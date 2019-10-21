Step 3: Code Hunting
====================

If you plan to submit a feature or bug fix to Eclcloud, you must be
able to prove your code correctly works with the Enterprise Cloud
service in question.

One way of verifying this is through the [Enterprise Cloud API reference
documentation](https://ecl.ntt.com/en/documents/api-references/).
However, the API docs might either be incorrect or they might not provide all of
the details we need to know in order to ensure this field is added correctly.

Code Hunting Tips
-----------------

Enterprise Cloud services differ from one to another. Code is organized in different
ways. However, the following tips should be useful across all projects.

* The logic which implements Create and Delete actions is usually either located
  in the "model" or "controller" portion of the code.

* Use Github's search box to search for the exact field you're working on.
  Review all results to gain a good understanding of everywhere the field is
  used.

* When adding a field, look for an object model or a schema of some sort.

---

Proceed to [Step 4](step-04-pull-requests.md) to learn about Acceptance
Testing.
