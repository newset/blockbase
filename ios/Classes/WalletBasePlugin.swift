import Foundation
import Flutter

public class SwiftPluginPocPlugin:NSObject{
    public static func register(with registrar: FlutterPluginRegistrar) {
        let channel = FlutterMethodChannel(name: "swift_plugin_poc", binaryMessenger: registrar.messenger())
        let instance:SwiftPluginPocPlugin  = SwiftPluginPocPlugin()
        if instance is FlutterPlugin{
            registrar.addMethodCallDelegate(instance as! FlutterPlugin, channel:channel)
        } else {
            print("THE PLUGIN DOES NOT IMPLEMENTIG FlutterPlugin")
        }
    }
    
    public func handle( _ call:FlutterMethodCall, result: FlutterResult ){
        if call.method == "getPlatformVersion" {
            result("iOS response : \(UIDevice.current.systemVersion)")
        } else {
            result( FlutterMethodNotImplemented)
        }
    }

    public func dummy () {
        version();
    }
}
