Project

Goal: Migrate all legacy pdlc language generators to the Go-based AST pipeline.
Repo root: /home/kapablanka/repos/pdl
Environments: go1.25.0 (linux/amd64), pnpm 8+, clang toolchain (legacy C++ reference)

Current Status

Phase: execution
Next step id: S4

S1:

-Write generator for js based on ast
-check parity with c++ generator
-mark as done

S1.1:
-Repeat S1 for PHP and C#

S2:

-Write generator for java based on ast
-compile generate and make sure the output has no syntax error
-mark as done

-Repeat S2 for kotlin, rust, c++ and python
  - Kotlin, Rust and C++ generators now emit AST-driven sources (data classes, structs, headers). Python remains outstanding.

Optimizations:

-Introduce a bounded worker-pool for AST generator fanout so large projects parallelize `pdlgen` runs without overwhelming the host.


S3: db2pdl - Java, Kotiln
-Java, koltin generator
-make them have same functionality as GO and PHP
-fluent queries, type safe


S4:  db2pdl - C++, Rust
-C++, Rust generator
-make them have same functionality as GO and PHP
-fluent queries, type safe

S5:  db2pdl - python
-python generator
-make them have same functionality as GO and PHP
-fluent queries, type safe