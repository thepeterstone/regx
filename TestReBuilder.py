#!/usr/bin/env python

import unittest
from ReBuilder import ReBuilder

class TestReBuilder(unittest.TestCase):
  def testReBuilderMatchesExact(self):
    builder = ReBuilder()
    result = builder.addPattern('a');
    assert(result.rule().match('a'))

  def testIncrementalAddPattern(self):
    builder = ReBuilder();
    builder.addPattern('a').addPattern('b');
    assert(builder.rule().match('a'))
    assert(builder.rule().match('b'))

if __name__=='__main__':
  unittest.main()
