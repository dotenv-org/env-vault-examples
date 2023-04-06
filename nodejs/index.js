const crypto = require('crypto')

function encrypt(key, message) {
  // set up key and nonce
  key = _decodeKey(key)
  let nonce = _generateNonce()

  console.log('key', key.length, key)
  console.log('nonce', nonce.length, nonce)

  // set up cipher
  let cipher = crypto.createCipheriv('aes-256-gcm', key, nonce)

  // generate ciphertext
  let ciphertext = ''
  ciphertext += cipher.update(message, 'utf8', 'hex')
  ciphertext += cipher.final('hex')
  ciphertext += cipher.getAuthTag().toString('hex')

  // prepend nonce
  ciphertext = nonce.toString('hex') + ciphertext

  // base64 encode output
  return Buffer.from(ciphertext, 'hex').toString('base64')
}

function decrypt(key, ciphertext) {
  // setup key
  key = _decodeKey(key)

  // base64 decode input
  ciphertext = Buffer.from(ciphertext, 'base64')

  // extract nonce
  const nonce = ciphertext.slice(0, 12)

  // extract authtag
  const authTag = ciphertext.slice(-16)

  // extract ciphertext
  ciphertext = ciphertext.slice(12, -16)

  // set up cipher
  const cipher = crypto.createDecipheriv('aes-256-gcm', key, nonce)
  cipher.setAuthTag(authTag)

  let message = ''
  message += cipher.update(ciphertext)
  message += cipher.final()

  return message
}

function _decodeKey(key) {
  return Buffer.from(key, 'hex')
}

function _generateNonce() {
  return crypto.randomBytes(_nonceBytes())
}

function _keyBytes() {
  return 32
}

function _authTagBytes() {
  return 16
}

function _nonceBytes() {
  return 12
}

const EnvVaultExample = {
  encrypt,
  decrypt
}

module.exports = EnvVaultExample

let key = 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'
let message = 'HELLO'

let ciphertext = EnvVaultExample.encrypt(key, message)
console.log(ciphertext)

let decryptedtext = EnvVaultExample.decrypt(key, ciphertext)
console.log(decryptedtext)
