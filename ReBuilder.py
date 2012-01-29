#!/usr/bin/env python
import re

class ReBuilder:
  atoms = []
  concatenator = ')('

  def addLine(self, line):
    if self.atoms == []:
      self.atoms = self.condense(self.split(line))
    if not self.match(line):
      self.atoms = self.merge(self.atoms, self.split(line))
    return self

  def rule(self):
    return re.compile('(' + self.concatenator.join(self.atoms) + ')')

  def merge(self, orig, addenda):
    return self.condense(orig + addenda)
  
  def match(self, test):
    #if not self.rule().match(test): print self.rule().pattern, test
    return self.rule().match(test)

  def split(self, pattern):
    bits = pattern.split(self.concatenator)
    output = []
    for i in range(len(bits) - 1):
      mutator = AtomMutate()
      mutator.addPattern(bits[i])
      if len(self.atoms) > i: mutator.addPattern(self.atoms[i])
      output += mutator.rule().pattern
    return output 


  def condense(self, atoms):
    return atoms


class AtomMutate(ReBuilder):
  concatenator = '|'

  def addPattern(self, pattern):
    if self.atoms == []:
      self.atoms = self.condense(self.split(pattern))
    else:
      if not self.match(pattern):
        self.atoms = self.merge(self.atoms, self.split(pattern))
    return self

  def condense(self, atoms):
    digits = alpha = xdigits = 0
    for symbol in atoms:
        if re.match('\d', symbol): digits += 1
        if re.match('[A-Za-z]', symbol): alpha += 1
        if re.match('[0-9a-fA-F]', symbol): xdigits += 1

    #if len(atoms) > 3: print "D: %d\tX: %d\tA: %d\t(%s)" % (digits, xdigits, alpha, atoms)
    if digits > 3 and alpha == 0:
        return [ '\d' ]

    if alpha > 5 and digits == 0:
        return [ '[A-z]' ]

    if xdigits == alpha + digits and xdigits > 4:
        return [ '[0-9a-fA-F]' ]

    return atoms

  def split(self, pattern):
      return list(pattern)


if __name__=='__main__': 
    import TestReBuilder, unittest
    unittest.main(TestReBuilder)
