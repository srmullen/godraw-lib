godraw-lib
==========

Drawing library.

TODO
----

- Implement Layers


Refactor
--------

- Simplify Line and Point with Vector and Coords

- Create Axi Layers and provide the ability to switch between layers. A Layer is a collection of items that will be drawn together.
  Can switch between layers when producing code, but when printing each layer will be printed all together regardless of the order that shapes were added to layers. All calls to the underlying svg renderer should be defered until Done() is called.