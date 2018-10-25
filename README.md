# tgen

`tgen` is an extremely simple templated find/replace tool.

It was created to solve for replacing templated values in config files.
 
Replacements can be described via the CLI, or in a substitutions file. 

## Usage

Given the following templated file, `template`:
```
Hello,
The secret is {{ SECRET }}.
```
And the substitution file, `substitutions`:
```
SECRET=I am a Cylon
```
The following command would execute as follows:
```
$ tgen -t template -f substitutions
Hello,
The secret is I am a Cylon.
```
You may also pass substitutions via the command line:
```
$ tgen -t template -r SECRET=I am a Cylon.
Hello,
The secret is I am a Cylon.
```

That's really all there is to it! 

