# Renovate Configuration Guide

## Overview

This project uses **Renovate** for automated dependency updates with automerge enabled to reduce notification emails.

## Configuration Details

### Automerge Settings

- **Enabled**: Yes ✅
- **Strategy**: Squash commits
- **Scope**: Minor and patch updates automatically merged
- **Major Updates**: Require manual approval

### Update Rules

#### Go Updates
```json
{
  "matchDatasources": ["golang-version"],
  "matchUpdateTypes": ["minor", "patch"],
  "automerge": true
}
```
- Automatically merges minor and patch Go version updates
- Requires approval for major versions
- Minimum 3 days before auto-merging

#### Dependencies
- **Patch Updates**: Auto-merge ✅
- **Minor Updates**: Auto-merge ✅
- **Major Updates**: Manual approval required ⚠️

#### Docker Images
- **Patch Updates**: Auto-merge ✅
- **Minor Updates**: Manual approval ⚠️
- **Major Updates**: Manual approval ⚠️
- Minimum 5 days before auto-merging

#### Vulnerability Alerts
- **All Severity**: Auto-merge ✅
- **Security Priority**: Highest

### Notification Policy

```
📧 Email Notifications: DISABLED
✅ Auto-merge Active: Enabled
🔄 Concurrent PRs: Max 3
📊 Daily PR Limit: 5
```

**Result**: You won't get emails for:
- Patch updates (automatically merged)
- Minor updates (automatically merged)
- Non-security issues

**You WILL get emails for**:
- Major version updates (manual review needed)
- Failing builds
- Conflicts requiring attention

### Commit Conventions

All Renovate commits follow semantic versioning:
```
chore(deps): update go to 1.21.2
chore(deps): update module github.com/package
chore(deps): update postgres image
```

### Safety Features

1. **Tests Required**: All updates run tests before merging
2. **Staged Rollout**: New versions wait 3-5 days
3. **Conflict Detection**: Alerts on merge conflicts
4. **Dependency Verification**: Checks compatibility

### Configuration File Location

- **File**: `renovate.json`
- **Root**: Repository root
- **Format**: JSON with schema validation

## How to Modify

### Enable Email Notifications
```json
{
  "prCreationLimit": 0,
  "automerge": false
}
```

### Disable Automerge for Specific Package
```json
{
  "packageRules": [
    {
      "matchPackagePatterns": ["specific-package"],
      "automerge": false
    }
  ]
}
```

### Change Merge Strategy
```json
{
  "automergeType": "merge",  // or "squash", "rebase"
  "automergeStrategy": "merge"
}
```

## Monitoring Automerged Updates

### Check Recent Merges
```bash
git log --grep="chore(deps)" --oneline | head -20
```

### View Renovate Dashboard
- Visit: `https://app.renovatebot.com/dashboard`
- Select your repository
- See all update history

### Disable Renovate Temporarily
Comment in PR:
```
@renovatebot pause
```

Resume:
```
@renovatebot resume
```

## Troubleshooting

### Too Many PRs?
Reduce `prCreationLimit`:
```json
{
  "prCreationLimit": 2
}
```

### Want to Review Before Merge?
Disable automerge:
```json
{
  "automerge": false
}
```

### Merge Conflicts?
Renovate will:
1. Detect conflict
2. Create PR but not merge
3. Await manual resolution
4. You fix and approve

### Update Too Fast?
Increase minimum age:
```json
{
  "minimumReleaseAge": "7 days"
}
```

## Best Practices

✅ **Do**:
- Keep automerge enabled for stability
- Review major updates manually
- Check the dependency dashboard weekly
- Archive old Renovate PRs

❌ **Don't**:
- Disable all notifications
- Auto-approve security updates without review
- Merge conflicting updates blindly
- Ignore major version warnings

## Renovate Commands

### Via Comments on PR
```
@renovatebot rebase
@renovatebot recreate
@renovatebot squash
@renovatebot pause
@renovatebot resume
@renovatebot bump
@renovatebot unschedule
```

## Resources

- [Renovate Docs](https://docs.renovatebot.com/)
- [Go Dependencies](https://docs.renovatebot.com/modules/datasource/go/)
- [Docker Support](https://docs.renovatebot.com/modules/datasource/docker/)
- [Configuration Reference](https://docs.renovatebot.com/configuration-options/)

---

**Last Updated**: May 30, 2026  
**Status**: Active & Auto-merging  
**Email Notifications**: Disabled for routine updates
