# Exercism Go exercises

## Reimplementing an exercise from scratch

When the user wants to redo an already-submitted exercise from scratch,
follow these steps in order:

1. Tell the user to go to the exercise's page on exercism.org and click
   "Update to the latest version."

2. Wait for the user to confirm they've done it. Do not run the download
   command before this confirmation.

3. Once confirmed, run:
   ```bash
   exercism download --exercise=<slug> --track=go --force
   ```

4. Replace the implementation file (e.g. `simple_linked_list.go`) with the
   exercise file *without* the implementation, fetched from GitHub:
   ```bash
   curl -sf "https://raw.githubusercontent.com/exercism/go/main/exercises/practice/<slug>/<slug>.go" \
     -o <path-to-local-exercise-dir>/<slug>.go
   ```
   This step is necessary because `exercism download` always restores the
   user's latest submitted solution for that file, never a blank stub, once
   an iteration has been submitted.
