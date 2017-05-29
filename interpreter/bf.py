import sys

class Bf(object):
    def __init__(self, prog):
        self.prog = prog

    def run():
        pass

def main():
    if len(sys.argv) < 2:
        print "specify the file to run"
        return

    with open(sys.argv[1]) as f:
        bf = Bf(f.read())
        bf.run()
