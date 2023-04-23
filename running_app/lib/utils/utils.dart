import 'dart:convert';
import 'dart:developer';

import 'package:encrypt/encrypt.dart';

class EncryptData {
  static Encrypted? encrypted;
  static var decrypted;

  String encryptAES(String plainText, keyStr) {
    final key = Key.fromBase16(keyStr);
    final iv = IV.fromSecureRandom(16);
    final encrypter = Encrypter(AES(key));
    //final encrypter = Encrypter(AES(key, mode: AESMode.cfb64));

    encrypted = encrypter.encrypt(plainText, iv: iv);

    log("URI: ${Uri.encodeComponent(encrypted!.base64)}");
    //return (Uri.encodeComponent(encrypted!.base64));

    return (base64.encode(iv.bytes + encrypted!.bytes));
  }

  String decryptAES(String plainText, keyStr) {
    final key = Key.fromUtf8(keyStr);
    final iv = IV.fromLength(16);
    final encrypter = Encrypter(AES(key));
    decrypted = encrypter.decrypt(encrypted!, iv: iv);
    return (decrypted);
  }
}
