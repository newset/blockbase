use keys::Private;
use std::ffi::{CStr, CString};
use std::os::raw::c_char;

#[no_mangle]
pub extern "C" fn rust_greeting(to: *const c_char) -> *mut c_char {
    let c_str = unsafe { CStr::from_ptr(to) };
    let recipient = match c_str.to_str() {
        Err(_) => "there",
        Ok(string) => string,
    };

    CString::new("Hello ".to_owned() + recipient)
        .unwrap()
        .into_raw()
}

#[no_mangle]
pub extern "C" fn signTron() -> *mut c_char {
    let raw = "helloworld".as_bytes();
    let priv_key: Private = "04811f1b4c96b2f26d0ec6cc74a51386c62b4633c28bbb20a1f2a0b64e9368ff"
        .parse()
        .unwrap();

    let sign = priv_key.sign_digest(&raw);

    // yep, the magic
    sign[64] += 27;

    eprintln!("sign: {:?}", hex::encode(sign));

    // Ok(sign.as_bytes())
    return CString::new("123").unwrap().into_raw();
}

#[no_mangle]
pub extern "C" fn rust_cstr_free(s: *mut c_char) {
    unsafe {
        if s.is_null() {
            return;
        }
        CString::from_raw(s)
    };
}
