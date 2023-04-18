package main

import "log"

// implement
func encrypt(key string, message string) {
  log.Println(key)
  log.Println(message)
}

// implement
func decrypt(key string, ciphertext string) {
  log.Println(key)
  log.Println(ciphertext)
}

func main() {
  key := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
  message := "HELLO"

  encrypt(key, message)
  decrypt(key, "ciphertext")
  // ciphertext = EnvVaultExample.new.encrypt(key, message)
  //puts ciphertext

  // decryptedtext = EnvVaultExample.new.decrypt(key, ciphertext)
  //puts decryptedtext
}
