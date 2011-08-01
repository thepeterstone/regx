#!/usr/bin/env python
import re

class ReBuilder:
  atoms = []

  def addPattern(self, pattern):
    if self.atoms == []:
      self.atoms = self.split(pattern)
    else:
      if self.match(pattern):
        self.atoms = self.merge(self.atoms, self.split(pattern))
    return self

  def rule(self):
    return re.compile(self.atoms.join(''))

  def split(self, pattern):
    return pattern.split('')

  def merge(self, orig, addenda):
    return [orig, addenda]
  
  def match(self, test):
    return self.rule().match(test)



