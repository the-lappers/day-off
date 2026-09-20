# Lessons learned

Add an entry when something cost more than an hour, broke twice, or surprised someone.
Newest first. Keep each entry to four lines. This feeds the final report sections on
limitations, adaptations, and risk.

Template:

```
## YYYY-MM-DD — short title
What happened:
Root cause:
Fix / what we do now:
```

## 2026-09-20 — gitleaks scanned only one branch
What happened: The pre-flight secret scan reported 3 commits when the repo had more.
Root cause: `gitleaks git` scans the checked-out branch unless told otherwise.
Fix / what we do now: `make secrets` passes `--log-opts="--all"` so every ref is scanned.
