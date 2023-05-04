# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(blockbase_library VERSION 0.0.1 LANGUAGES C)

add_library(blockbase SHARED
  "blockbase.c"
)

set_target_properties(blockbase PROPERTIES
  PUBLIC_HEADER blockbase.h
  OUTPUT_NAME "blockbase"
)

target_compile_definitions(blockbase PUBLIC DART_SHARED_LIB)
