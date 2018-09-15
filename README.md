# tgen

tgen is an extremely simple template generator

A simple command line utility for replacing values in a template. 

It was created to work for templates where the only thing needed is a simple find and replace. 
Replacements can be described via the CLI, or in a substitutions file. 

## Usage

Given a templated file `template`:
```
Hello,
The secret is {{ SECRET }}.
```
And the substitution file `substitutions`:
```
SECRET=I am a Cylon
```
The following command would execute as follows:
```
$ tgen -t template -f substitutions
Hello,
The secret is I am a Cylon.
```



