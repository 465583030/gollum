TemplateJSON
============

TemplateJSON is a formatter that evaluates a text template with an input of a JSON message.


Parameters
----------

**TemplateJSONFormatter**
  TemplateJSONFormatter formatter that will be applied before the field is templated.
  Set to format.Forward by default.

**TemplateJSONTemplate**
  TemplateJSONTemplate defines the template to execute with text/template.
  This value is empty by default.
  If the template fails to execute the output of TemplateJSONFormatter is returned.

Example
-------

.. code-block:: yaml

	- "stream.Broadcast":
	    Formatter: "format.TemplateJSON"
	    TemplateJSONFormatter: "format.Forward"
	    TemplateJSONTemplate: ""
