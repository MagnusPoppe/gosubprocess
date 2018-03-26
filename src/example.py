if __name__ == '__main__':
    try:
        print("Ready!")
        while True:
            message = input("Waiting for input")
            print(message)
    except KeyboardInterrupt:
        print("Finished by keyboard!")
        exit(0)
    print("Finished for real!!")
