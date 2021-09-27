# go-netapp
golang's netapp client

## example

```golang

c, err := netapp.NewClient(
    <your endpoint>,
    <your version>,
    &netapp.ClientOptions{
    },
)
if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
}

qRes, _, err := c.QuotaReport.Report(&netapp.QuotaReportOptions{
        MaxRecords: 1,
        Query: &netapp.QuotaReportEntryQuery{
                QuotaReportEntry: &netapp.QuotaReportEntry{
                        QuotaTarget: quotaTarget,
                },
        },
})
```

## Contribution

1. Fork ([https://github.com/pepabo/go-netapp/fork](https://github.com/pepabo/go-netapp/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[pyama86](https://github.com/pyama86)
