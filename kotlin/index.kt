/* implement */
fun encrypt(key string, message string) {

}

fun decrypt(key string, ciphertext string) {

}

fun run() {
  key := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
  message := "HELLO"

  ciphertext = encrypt(key, message)
  message2 = decrypt(key, ciphertext)

  /* message 2 should equal message */
}
