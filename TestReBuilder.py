#!/usr/bin/env python

import unittest
from ReBuilder import *

class TestAtomMutate(unittest.TestCase):
  def testAtomMutateMatchesExact(self):
    builder = AtomMutate()
    result = builder.addPattern('a')
    assert(result.match('a'))

  def testIncrementalAddPattern(self):
    builder = AtomMutate()
    builder.addPattern('a').addPattern('b')
    assert(builder.match('a'))
    assert(builder.match('b'))

  def testCondenseFourDigits(self):
    builder = AtomMutate()
    builder.addPattern([ '1', '2', '3' ]);
    assert(not builder.match('5'))
    builder.addPattern('4')
    assert(builder.match('1'))
    assert(builder.match('5'))

  def testCondenseHexChars(self):
    builder = AtomMutate()
    builder.addPattern([ '1', 'a', '4', 'f', 'e', '6' ])
    assert(builder.match('1'))
    assert(builder.match('b'))
    assert(not builder.match('g'))

  def testCondenseHexWithDigits(self):
    builder = AtomMutate()
    builder.addPattern([ '1', '2', '3', '4', '5', 'a', 'b', 'f' ])
    assert(builder.match('d'))
    assert(builder.match('9'))
    assert(not builder.match('g'))

class TestReBuilder(unittest.TestCase):
  def testSingleChar(self):
    builder = ReBuilder()
    builder.addLine('a')
    assert(builder.match('a'))

  def testSingleLine(self):
    builder = ReBuilder()
    builder.addLine('This is a test line')
    assert(builder.match('This is a test line'))

  def testIncrementalAddLine(self):
    builder = ReBuilder()
    builder.addLine('a').addLine('b')
    assert(builder.match('a'))
    assert(builder.match('b'))

  def testCondenseDigits(self):
    builder = ReBuilder()
    builder.addLine('1')
    builder.addLine('2')
    builder.addLine('3')
    builder.addLine('4')
    assert(builder.match('5'))

  def testDateStamp(self):
    builder = ReBuilder()
    stamps = [ '2011-01-01', '2010-09-27', '2011-04-23', '2010-12-31', '2009-06-14' ]
    for stamp in stamps:
      builder.addLine(stamp)
      assert(builder.match(stamp))
    assert(builder.match('2011-03-12'))



if __name__=='__main__': unittest.main()
