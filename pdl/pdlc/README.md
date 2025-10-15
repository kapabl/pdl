# Building `pdlc`

All C++ sources, vcpkg metadata, and CMake files live inside this directory, so you can work entirely from here.

```bash
# from pdl/pdlc
./build.sh
```

The script clones and bootstraps vcpkg (if missing), installs the dependencies declared in `vcpkg.json`, configures CMake with the local vcpkg toolchain, and builds the `pdlc` binary into `build/pdlc/`, copying the executable to `pdl/bin/`.

Feel free to pass any usual CMake flags (e.g. `-DCMAKE_BUILD_TYPE=Release`, `-DCMAKE_TOOLCHAIN_FILE=...`) to the first step when you need a custom setup.
