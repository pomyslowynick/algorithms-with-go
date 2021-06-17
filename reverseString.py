# Write a function that takes a string and returns that string in reverse
# “Hello, World!” => “!dlroW ,olleH”

# "Hello"
def string_reverse(arg_string):
    result = ""
    
    if type(arg_string) != str:
        raise(Exception("Not a string"))
        
    if arg_string == "":
        return ""
        
    y = len(arg_string) - 1
    while y >= 0:
        result += arg_string[y]
        y -= 1
    
    return result

print(string_reverse('Hello'))
# @Test
# def testThatReverseWorks()

# @Test
# def testThatErrorIsRaise()
    