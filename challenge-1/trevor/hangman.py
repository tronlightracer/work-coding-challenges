from string import ascii_lowercase

def missing_chars(filename):
    f = open(filename, "r")
    read_letters, ascii = set(f.read()), set(ascii_lowercase)
    return "".join(ascii - read_letters)

print(missing_chars("hangman_text_file.txt"))

