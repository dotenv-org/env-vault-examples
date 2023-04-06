require "openssl"
require "securerandom"
require "base64"

class EnvVaultExample
  # Example
  # key: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  # message: "HELLO=world"
  #
  def encrypt(key, message)
    # set up key and nonce
    key = decode_key(key)
    nonce = generate_nonce

    # set up cipher
    cipher = OpenSSL::Cipher.new("aes-256-gcm")
    cipher.encrypt
    cipher.key = key
    cipher.iv = nonce
    cipher.auth_data = ""

    # generate ciphertext
    ciphertext = String.new
    ciphertext << cipher.update(message)
    ciphertext << cipher.final
    ciphertext << cipher.auth_tag

    # prepend nonce
    ciphertext = nonce + ciphertext

    # base64 encode output
    Base64.strict_encode64(ciphertext)
  end

  def decrypt(key, ciphertext)
    # set up key
    key = decode_key(key)

    # base64 decode input
    ciphertext = Base64.decode64(ciphertext)

    # extract nonce
    nonce = ciphertext.slice(0, nonce_bytes)
    ciphertext = ciphertext.slice(nonce_bytes..-1)

    # extra authtag
    auth_tag = ciphertext.slice(-auth_tag_bytes..-1)
    ciphertext = ciphertext.slice(0, ciphertext.bytesize - auth_tag_bytes)

    # set up cipher
    cipher = OpenSSL::Cipher.new("aes-256-gcm")
    cipher.decrypt
    cipher.key = key
    cipher.iv = nonce
    cipher.auth_tag = auth_tag
    cipher.auth_data = ""

    # decipher message
    message = String.new
    message << cipher.update(ciphertext)
    message << cipher.final

    message
  end

  private

  def decode_key(key)
    if key.encoding != Encoding::BINARY && key =~ /\A[0-9a-f]{#{key_bytes * 2}}\z/i
      key = [key].pack("H*")
    end

    raise ArgumentError, "Key must be #{key_bytes} bytes (#{key_bytes * 2} hex digits)" unless key && key.bytesize == key_bytes
    raise ArgumentError, "Key must use binary encoding" unless key.encoding == Encoding::BINARY

    key
  end

  def key_bytes
    32
  end

  def generate_nonce
    SecureRandom.random_bytes(nonce_bytes)
  end

  def nonce_bytes
    12
  end

  def auth_tag_bytes
    16
  end
end

key = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" 
message = "HELLO"

ciphertext = EnvVaultExample.new.encrypt(key, message)
puts ciphertext
decryptedtext = EnvVaultExample.new.decrypt(key, ciphertext)
puts decryptedtext
