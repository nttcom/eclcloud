Step 2: Create an Issue
=======================

Every patch / Pull Request requires a corresponding issue. If you're fixing
a bug for an existing issue, then there's no need to create a new issue.

However, if no prior issue exists, you must create an issue.

Reporting a Bug
---------------

When reporting a bug, please try to provide as much information as you
can.

Feature Request
---------------

If you've noticed that a feature is missing from Eclcloud, you'll also
need to create an issue before doing any work. This is start a discussion about
whether or not the feature should be included in Eclcloud. We don't want to
want to see you put in hours of work only to learn that the feature is out of
scope of the project.

Feature requests can come in different forms:

### Adding a Feature to Eclcloud Core

The "core" of Eclcloud is the code which supports API requests and
responses: pagination, error handling, building request bodies, and parsing
response bodies are all examples of core code.

Modifications to core will usually have the most amount of discussion than
other requests since a change to core will affect _all_ of Eclcloud.

### Adding a Missing Field

If you've found a missing field in an existing struct, submit an issue to
request having it added. These kinds of issues are pretty easy to report
and resolve.

You should also provide a link to the actual service's Python code which
defines the missing field.

There's one situation which can make adding fields more difficult: if the field
is part of an API extension rather than the base API itself.

### Adding a Missing API Call

If you've found a missing API action, create an issue with details of
the action.

You'll want to make sure the API call is part of the upstream Enterprise Cloud
project and not an extension created by a third-party or vendor. Eclcloud only
supports the Enterprise Cloud projects proper.

### Adding a Missing API Suite

Adding support to a missing suite of API calls will require more than one Pull
Request. However, you can use a single issue for all PRs.

Note how the issue breaks down the implementation by request types (Create,
Update, Delete, Get, List).

Also note how these issues provide links to the service's Python code. These
links are not required for _issues_, but it's usually a good idea to provide
them, anyway. These links _are required_ for PRs and that will be covered in
detail in a later step of this tutorial.

### Adding a Missing Enterprise Cloud service

These kinds of feature additions are large undertakings. Adding support for
an entire Enterprise Cloud project is something the Eclcloud team very much
appreciates, but you should be prepared for several weeks of work and
interaction with the Eclcloud team.

---

With all of the above in mind, proceed to [Step 3](step-03-code-hunting.md) to
learn about Code Hunting.
