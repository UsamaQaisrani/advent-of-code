from sys import argv
import os

def main():
    if len(argv) < 4:
        print("Format: python create_skeleton <yyyy> <number-of-days> <file-extension>")
    assert len(argv) == 4, "Format: python create_skeleton <yyyy> <number-of-days>"

    year = argv[1]
    assert len(year) == 4, "Year Format: YYYY"

    numDays = argv[2]
    assert numDays != "0", "Number of days should be more than 0."

    fileExtension = argv[3]
    assert fileExtension != "", "Please enter a file extension. Example: go"
    assert "." not in fileExtension, "Please enter file extenstion without '.'"

    parentDirPath = f"./{year}"
    if not os.path.exists(parentDirPath):
        os.makedirs(parentDirPath)

    for i in range(1,int(numDays)+1):
        if i < 10:
            childSubPath = f"day0{i}"
        else:
            childSubPath = f"day{i}"
        currChildPath = os.path.join(parentDirPath, childSubPath)
        if not os.path.exists(currChildPath):
            os.makedirs(currChildPath)

        mainFilePath = os.path.join(currChildPath, f"main.{fileExtension}")
        if not os.path.exists(mainFilePath):
            with open(mainFilePath, "w") as mainFile:
                mainFile.write("Start writing the code programmer!")

        inputFilePath = os.path.join(currChildPath, f"input.txt")
        if not os.path.exists(inputFilePath):
            with open(inputFilePath, "w") as inputFile:
                inputFile.write("")

        testInputFilePath = os.path.join(currChildPath, f"test.txt")
        if not os.path.exists(testInputFilePath):
            with open(testInputFilePath, "w") as testInputFile:
                testInputFile.write("")


if __name__ == "__main__":
    main()
