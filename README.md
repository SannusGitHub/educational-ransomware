# educational-malware
## ATTENTION
THIS REPOSITORY AND SUBSEQUENTLY SOFTWARE WAS CREATED FOR EDUCATIONAL AND ARCHIVAL PURPOSES ONLY! THE CODE IN THIS PROJECT HAS BEEN EDITED AND CHANGED FROM THE ACTUAL PROJECT.

STRICTLY PROHIBITED FOR ANY ILLEGAL USE, AND THE AUTHOR BEARS NO RESPONSIBILITY FOR ANY MISUSE.

FOR SAFETY AND EASE-OF-USE, THIS PROJECT HAS BEEN MODIFIED TO DO THE FOLLOWING:
* ONLY CHANGE FILES WITH THE ".cool" EXTENSION
* PROVIDE THE DECRYPTION KEY IN PLAIN TEXT IN THE GENERATED TEXT FILE
* REMOVAL OF A BINDER WHICH COULD MERGE A CLEAN PROGRAM WITH THE MALICIOUS ONE
* REMOVAL OF THE BUNDLED CODE THAT GENERATES AN APPROPRIATE DECRYPTOR FOR THE END-USER

SIDE-NOTE: THIS PROJECT WILL MOST LIKELY NOT WORK ANYMORE, DUE TO THE FACT THAT IT'S SIGNATURE HAS BEEN GIVEN TO A DATABASE AND THUS IS NOW PROPERLY VULNERABLE.

## Overview
The goal of this project was to create a fully functioning ransomware program that encrypts the personal data stored on a machine, with the following requirements:
* Encrypts everything in a specific directory with a custom extension
* Provides a sample readme file with information in regards to how to get a key for decryption
* Provides a decrypting program which takes the key and decrypts the previously encrypted files

##

## Questions
### What is ransomware?
Ransomware is a specific type of malware which prevents an individual from accessing their device and/or the data stored on it (usually through encryption) until a ransom is paid

### How does this software attempt to bypass AV systems?
The program can be compiled with the "garble" package available for go to avoid further detection, using the command
`garble -literals -tiny -seed=random build encryptor.go`
and by also hardcoding some variable names to random strings (which is done in this project already)

### How does this ransomware work?
The ransomware uses a randomly generated AES256 key which is made upon runtime. The process of the program is as follows:
1. A random AES256 key is generated
2. The program then gets all of the files under the home (`Users`) directory
3. The program iterates recursively over all of the files in the home (`Users`) directory checking for files which are encryptable (in this case formats with `.cool` as an extension)
4. If a file is encryptable it gets encrypted with the random key and it's extension is changed
5. Upon encrypting all of the files it generates a contact info text file called `README_IMPORTANT.txt` on the users desktop
6. Upon running the decryptor.exe, you can provide it with the key appended to the `README_IMPORTANT.txt` file and decrypt files

### How does the decrypting program work?
The decrypting program works the same as the encryption, although just in reverse. The process of the program is as follows:
1. When ran, it asks for the users decryption key so it can start working on decrypting the files
2. It iterates and checks over all of the files stored in the home (`Users`) directory
3. When checking a file that's stored, it opens the contents of the file to read
4. It prepares a key and turns it into a format the program can use
5. It creates a "lock-breaking" tool using the AES algorithm
6. It grabs a small part from the beginning of the file, called a "nonce"
7. Using the "lock-breaking" AES tool and the nonce it tries to decrypt the content
8. Upon successfully checking whether the key is matching and decryption is successful it saves the unlocked file

##

## Usage
Either run the provided executable of `encryptor.exe` or run it by doing the following commands in a terminal:

```
garble -literals -tiny -seed=random build encryptor.go
go run encryptor.exe
```
...and the same with the subsequent decryptor.go file if needed

Upon successful execution, files in the home (`Users`) directory will be decrypted and a provided text file will be provided on the desktop. The text file (for the convenience of this project) contains the key needed to decrypt the encrypted files

Then run the `decryptor.exe`and input the decryption key to decrypt files