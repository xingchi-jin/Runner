# Github action Task
Github action task natively executes a github action plugin.

## Sample
Let's use a [github checkout action](https://github.com/marketplace/actions/checkout) as the example. You can find the definition of this action from [here](https://github.com/actions/checkout/blob/main/action.yml)

To use it in Runner, you can do it without modification. Here it is trying to checkout a branch.

To do it in Github Action, here is the [way](https://github.com/actions/checkout?tab=readme-ov-file#checkout-a-different-branch)

To do it in Harness Runner, below is the way
```
taskGroup
- name: checkout
  type: github_action
  spec:
    // here provide the same spec as Github Action.
    uses: actions/checkout@v4
    with:
      ref: my-branch
  exports:
    - name: <output name of the checkout action>
```