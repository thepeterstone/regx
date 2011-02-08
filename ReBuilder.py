#!/usr/bin/env python
import re

class ReBuilder:
  rule = None

  def addPattern(self, pattern):
    if (self.rule == None):
      self.rule = re.compile(pattern)
      return self
    if (self.rule.match(pattern)):
      return self
    self.rule = re.compile(self.rule.pattern + '|' + pattern)
    return self



