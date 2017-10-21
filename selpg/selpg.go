package main

import (
	"io"
	"os"
	"fmt"
	"strings"
	//"flag"
	"strconv"
	"bufio"
	//"os/exec"
	//"runtime"
)

type sp_args struct {
	StartPage int
	EndPage int
	InFilename string
	PageLen int
	PageType string
	PrintDest string
} 

const INBUFSIZ = 16*1024

var progname string

func usage() {
	fmt.Printf("\nUSAGE: %s -sstart_page -eend_page [-f | -llines_per_page] [-ddest] [in_filename]\n", progname)
}

func process_args(ac int, av []string, sa *sp_args) {
	var s1,s2 string
	var argno,i int
	// var sa = &psa

	if ac < 3 {
		fmt.Printf("%s: not enough arguments\n", progname)
		usage()
		os.Exit(1)
	}

	s1 = av[1]
	if strings.EqualFold("-s", s1[0:2]) != true {
		fmt.Printf("%s: 1st arg should be -sstart_page\n", progname)
		usage()
		os.Exit(2)
	}
	tp := s1[2:]
	i, _ = strconv.Atoi(tp)
	if i < 1 {
		fmt.Printf("%s: invalid start page %s\n", progname, tp)
		usage()
		os.Exit(3)
	}
	sa.StartPage = i

	s1 = av[2]
	if strings.EqualFold("-e", s1[0:2]) != true {
		fmt.Printf("%s: 2nd arg should be -eend page\n", progname)
		usage()
		os.Exit(4)
	}
	tp = s1[2:]
	i, _ = strconv.Atoi(tp)
	if i < 1 || i < sa.StartPage {
		fmt.Printf("%s: invalid end page %s\n", progname, tp)
		usage()
		os.Exit(5)
	}
	sa.EndPage = i

	argno = 3
	for argno <= (ac-1) && av[argno][0:1] == "-" {
		s1 = av[argno]
		switch s1[1:2] {
		case "l":
			s2 = s1[2:]
			i, _ = strconv.Atoi(s2)
			if i < 1  {
				fmt.Printf("%s: invalid page length %s\n", progname, s2)
				usage()
				os.Exit(6)
			}
			sa.PageLen = i
			argno++
		case "f":
			if s1 != "-f" {
				fmt.Printf("%s: option should be \"-f\"\n", progname)
				usage()
				os.Exit(7)
			}
			sa.PageType = "f"
			argno++
		case "d":
			s2 = s1[2:]
			if len(s2) < 1 {
				fmt.Printf("%s: -d option requires a printer destination\n",progname)
				usage()
				os.Exit(8)
			}
			sa.PrintDest = s2
			argno++
		default:
			fmt.Printf("%s: unknown option %s\n", progname, s1)
			usage()
			os.Exit(9)
		} // end switch
	} // end for

	if argno <= (ac -1) {
		sa.InFilename = av[argno]
		file, err := os.OpenFile(av[argno], os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("Unable to read from %s", av[argno])
				file.Close()
				os.Exit(10)
			}
		}
	 }

	 // 断言
}
func process_input(sa sp_args) {
	var fin, fout *os.File
	var LineCtr, PageCtr int
	var line, page string
	var err error
	/* var cmd	*exec.Cmd
	var s1 string */

	if sa.InFilename == "" {
		fin = os.Stdin
	} else {
		fin, err = os.Open(sa.InFilename)
		if err != nil {
			fmt.Printf("%s: could not open input file \"%s\"\n")
			os.Exit(11)
		}
	}

	//设置缓冲区

	rd := bufio.NewReader(fin) // new rd
	// wd := bufio.NewWriter(fout) new wd

	if sa.PrintDest == "" {
		fout = os.Stdout
	} /* else { //作为指定线程的输入
		// os.Stdout.Flush()
		s1 = fmt.Sprintf("lp -d%s", sa.PrintDest)
		cmd = exec.Command(s1)
		wd, err = cmd.StdinPipe()
		if err != nil {
			fmt.Printf("%s: could not open pipe to \"%s\"\n", progname, sa.PrintDest)
			os.Exit(12)
		}
		//defer fin.Close()
 		cmd.Start() // 需要 fout.Close()
	} */


	if sa.PageType == "l" {
		LineCtr = 0
		PageCtr = 1

		for  {
			line, err = rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			LineCtr++
			if LineCtr > sa.PageLen {
				PageCtr++
				LineCtr = 1
			}
			if PageCtr >= sa.StartPage && PageCtr <= sa.EndPage {
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else { // type = "t"
		PageCtr = 0
		for {
			page, err = rd.ReadString('\f')
			if err != nil || io.EOF == err {
				break
			}
			PageCtr++
			if PageCtr >= sa.StartPage && PageCtr <= sa.EndPage {
				fmt.Fprintf(fout, "%s", page)
			}
		}
	}

	//ending input
	fin.Close()
	// fout.Flush()
	if sa.PrintDest != "" {
		fout.Close()
	}
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}

func main()  {
	var sa sp_args
	progname = os.Args[0]

	sa.StartPage, sa.EndPage = -1, -1
	sa.InFilename = ""
	sa.PageLen = 72
	sa.PageType = "l"
	sa.PrintDest = ""

	process_args(len(os.Args), os.Args, &sa)
	process_input(sa)

	return 
}

