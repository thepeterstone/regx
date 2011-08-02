#!/usr/bin/env python
import re

class ReBuilder:
  atoms = []

  def addPattern(self, pattern):
    if self.atoms == []:
      self.atoms = self.split(pattern)
    else:
      if not self.match(pattern):
        self.atoms = self.merge(self.atoms, self.split(pattern))
    return self

  def rule(self):
    return re.compile('|'.join(self.atoms))

  def split(self, pattern):
    return list(pattern)

  def merge(self, orig, addenda):
    return orig + addenda
  
  def match(self, test):
    return self.rule().match(test)



