package main

import (
    "flag"
    "fmt"
    "os"
)

func usage() {
    fmt.Fprintf(os.Stderr, "check_mem v0.1 (k0d-nagios-plugins) v0.1\n")
    fmt.Fprintf(os.Stderr, "Project: https://github.com/k0d/nagios-plugins\n")
    fmt.Fprintf(os.Stderr, "Copyright (c) 2011 Mark Andrews <m@k0d.se>\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Nagios plugin to monitor memory usage\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Usage:\n")
    fmt.Fprintf(os.Stderr, "check_mem [-hVv]\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Options:\n")
    fmt.Fprintf(os.Stderr, " -h, --help\n")
    fmt.Fprintf(os.Stderr, "    Print detailed help screen\n")
    fmt.Fprintf(os.Stderr, " -V, --version\n")
    fmt.Fprintf(os.Stderr, "    Print version information\n")
    fmt.Fprintf(os.Stderr, " -v, --verbose\n")
    fmt.Fprintf(os.Stderr, "    Show details for command-line debugging (Nagios may truncate output)\n")
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Send email to nagios@k0d.se if you have questions, to submit patches or\n")
    fmt.Fprintf(os.Stderr, "suggest improvements.\n")
    os.Exit(3)
}

func version() {
    fmt.Fprintf(os.Stderr, "Nagios Plugin to monitor memory usage - check_mem v1.0\n")
    os.Exit(3)
}

var flagVersion *bool = flag.Bool("V", false, "Print version number")

func main() {
    flag.Usage = usage
    flag.Parse()

    if *flagVersion {
        fmt.Println("Example Nagios Plugin v1.0");
        os.Exit(3);
    }
    fmt.Printf("EXAMPLE OK: Working\n");
    os.Exit(0);
}