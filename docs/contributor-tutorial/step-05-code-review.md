Step 6: Code Review
===================

Once you've submitted a Pull Request, three things will happen automatically:

1. Travis-CI will run a set of simple tests:

    a. Unit Tests

    b. Code Formatting checks

    c. `go vet` checks

2. Coveralls will run a coverage test.

Depending on the results of the above, you might need to make additional
changes to your code.

While you're working on the finishing touches to your code, it is helpful
to add a `[wip]` tag to the title of your Pull Request.

Request a Code Review
---------------------

When you feel your Pull Request is ready for review, please leave a comment
requesting a code review. If you don't explicitly ask for a code review, a
core member might not know the Pull Request is ready for review.

Additionally, if there are parts of your implementation that you are unsure
about, please ask for help. We're more than happy to provide advice.

During the code review process, a core member will review the code you've
submitted and either request changes or request additional information.
Generally these requests fall under the following categories:

1. Code which needs to be reformatted (See our [Style Guide](/docs/STYLEGUIDE.md)
   for conventions used.

2. Requests for additional information about the validity of something. This
   might happen because the included supporting service code URLs don't have
   enough information.

3. Missing unit tests or acceptance tests.

Submitting Changes
------------------

If a code review requires changes to be submitted, please do not squash your
commits. Please only add new commits to the Pull Request. This is to help the
code reviewer see only the changes that were made.

It's Never Personal
-------------------

Code review is a healthy exercise where a new set of eyes can sometimes spot
items forgotten by the author.

Please don't take change requests personally. Our intention is to ensure the
code is correct before merging.

---

Once the code has been reviewed and approved, a core member will merge your
Pull Request.

Please proceed to [Step 6](step-06-congratulations.md).
