// selpg
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type selpg_args struct{
	start_page int
	end_page int 
	in_filename sting
	page_len int
	page_type int
	print_dest string
}

var progname string

func process_args(ac int, av []string, pas *selpg_args){
	s1 string
	s2 string
	argno int
	i int
	if (ac < 3){
		fmt.fpFprintf(os.stderr, "%s: not enough arguments\n",progname);
		usage()
		os.Exit(1)
	}
	
	s1 = av[1];
	if(s1[0] != '-' || s1[1] != 's'){
		fmt.fFprint(os.stderr, "%s: 1st arg should be -sstart_page\n", progname);
		usage();
		os.Exit(2);
	}
	i, err:=strconv.Atoi(s1[2:])
	if( i <1 || i > MAX_INT || err != nil){
		fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, s1[2:])
		usage()
		os.Exit(3)
	}
	
	pas.start_page = i;
	
	s1 = av[2]
	if(s1[0] != '-' || s1[1] != 'e'){
		fmt.fFprint(os.stderr, "%s: 2nd arg should be -eend_page\n", progname);
		usage();
		os.Exit(4);
	}
	i, err:=strconv.Atoi(s1[2:])
	if( i < psa.start_page || i <1 || i > MAX_INT || err != nil){
		fmt.Fprintf(os.Stderr, "%s: invalid end page %s\n", progname, s1[2:])
		usage()
		os.Exit(5)
	}
	psa.end_page = i
	
	argno = 3
	for argno <= (ac-1) && av[argno][0] == '-'{
	    s1 = av[argno]
		switch s1[2]{
			
			case 'l':
			s2 = s1[2:]
			i, err := strconv.Atoi(s2)
			if (err != nil) || (i < 1 || i > MAX_INT) {
				fmt.Fprintf(os.Stderr, "%s: invalid page length %s\n", progname, s1[2:])
				usage()
				os.Exit(6)
			}
			psa.page_len = i
			argno++
			
			case 'f':
			if s1[0:2] != "-f" {
				fmt.Fprintf(os.Stderr, "%s: option should be \"-f\"\n", progname)
				usage()
				os.Exit(7)
			}
			psa.page_type = 'f'
			argno++
			
			case "d":
			s2 = s1[2:]
			if len(s2) < 1 {
				fmt.Fprintf(os.Stderr,
					"%s: -d option requires a printer destination\n", progname)
				usage()
				os.Exit(8)
			}
			psa.print_dest = s2
			argnow++

		    default:
			fmt.Fprintf(os.Stderr, "%s: unknown option %s\n", progname, s1)
			usage()
			os.Exit(9)
		}
		if(argno <= (ac -1)){
			psa.in_filename = av[argno]
			
			_, err := os.Stat(psa.in_filename)
			
		    if os.IsNotExist(err) {
			fmt.Printf("%s: input file \"%s\" does not exist\n", progname, psa.in_filename)
			os.Exit(10)
		    }
		}
		
	}
}

func process_input(sa selpg_args){
	var fin *os.File
	var fout *os.File
	var s1 string
	var line_ctr int
	var page_ctr int
	
	if sa.in_filename[0] == '~' {
		fin = os.Stdin
	} else {
		fin, _ = os.Open(sa.in_filename)
		if fin == nil {
			fmt.Printf("%s: could not open input file \"%s\"\n", progname, sa.in_filename)
			os.Exit(12)
		}
	}

	if sa.print_dest[0] == '~' {
		fout = os.Stdout
	} else {

		s1 = sa.print_dest
		fout, _ = os.Open(s1)
		if fout == nil {
			fmt.Printf("%s: could not open pipe to \"%s\"\n", progname, s1)
			os.Exit(13)
		}
	}

	if sa.page_type == 'l' {
		line_ctr = 0
		page_ctr = 1
		rd := bufio.NewReader(fin)
		for {
			line, ere := rd.ReadString('\n')
			if ere != nil {
				break
			}
			line_ctr += 1
			if line_ctr > sa.page_len {
				page_ctr += 1
				line_ctr = 1
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
				_, _ = io.WriteString(fout, line)
			}

		}
	} else {
		page_ctr = 1
		rd := bufio.NewReader(fin)
		for {
			line, ere := rd.ReadString('\n')
			if ere != nil {
				break
			}
			for _, v := range line {
				if v == '\f' {
					page_ctr += 1
				}
				if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
					_, _ = io.WriteString(fout, string(v))
				}
			}
		}
	}

	if page_ctr < sa.start_page {
		fmt.Printf("%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Printf("%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.end_page, page_ctr)
	}
	fin.Close()
	fout.Close()
}

func usage() {
	fmt.Printf("\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ]\n [ -ddest ] [ in_filename ]\n", progname)
}

func main() {
	var sa selpg_args
	
	progname = os.Args[0]
	sa.start_page = -1
	sa.end_page = -1
	sa.in_filename = ""
	sa.pagepage_type = 'l'
	sa.prprint_dest[0] = ""

}
