Sequence
========

Sequence is a formatter that allows prefixing a message with the message's sequence number .


Parameters
----------

**SequenceSeparator**
  SequenceSeparator sets the separator character placed after the sequence number.
  This is set to ":" by default.

**SequenceDataFormatter**
  SequenceDataFormatter defines the formatter for the data transferred as message.
  By default this is set to "format.Forward" .

Example
-------

.. code-block:: yaml

	- "stream.Broadcast":
	    Formatter: "format.Sequence"
	    SequenceFormatter: "format.Envelope"
	    SequenceSeparator: ":"
