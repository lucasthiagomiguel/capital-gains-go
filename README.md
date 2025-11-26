# Capital Gains CLI (Nubank Challenge)

This project implements a CLI that calculates taxes on capital gains from stock operations
according to the challenge spec.

## Run
Build:
```
make build
```

Example:
```
echo '[{"operation":"buy","unit-cost":10,"quantity":10000},{"operation":"sell","unit-cost":20,"quantity":5000}]' | ./capital-gains
```

## Tests
```
make test
```

## Notes
- Reads input line-by-line from stdin, each line is a JSON array of operations.
- Outputs JSON array of tax objects to stdout, one line per input line.
