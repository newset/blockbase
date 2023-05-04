// You have generated a new plugin project without specifying the `--platforms`
// flag. An FFI plugin project that supports no platforms is generated.
// To add platforms, run `flutter create -t plugin_ffi --platforms <platforms> .`
// in this directory. You can also find a detailed instruction on how to
// add platforms in the `pubspec.yaml` at
// https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:ffi';
import 'dart:io';

const String _libName = 'blockbase';

/// The dynamic library in which the symbols for [BlockbaseBindings] can be found.
final DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    return DynamicLibrary.open('$_libName.framework/$_libName');
  }
  if (Platform.isAndroid || Platform.isLinux) {
    return DynamicLibrary.open('lib$_libName.so');
  }
  if (Platform.isWindows) {
    return DynamicLibrary.open('$_libName.dll');
  }
  throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
}();

