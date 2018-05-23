# envless-listen

This is to demonstrate that on Windows, the SYSTEMROOT variable needs to be set in order for a `net.Listener` to work. This is especially not the case when executed from an os.Cmd with a custom environment.

## Setup

`cd listener && go build`
`go build`

## Working case

`./envless-listen -working`

## Failing case

`./envless-listen -failing`

Output:

`listen tcp 127.0.0.1:12345: socket: The requested service provider could not be loaded or initialized.`