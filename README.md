# 🌿 Plant Care Tracker

Some plants need water every few days, others only once a week.

This small Go application keeps a collection of houseplants, checks which ones need watering, and records basic care history.

---

## Console Preview

```
Today's Plant Checklist

✓ Monstera
Needs watering

✓ Aloe Vera
Watered yesterday

✓ Orchid
Needs watering

----------------------

Plants: 3
Need watering: 2
History records: 5
```

---

## Project Overview

| File | Purpose |
|------|---------|
| plant.go | Plant model |
| garden.go | Collection of plants |
| reminder.go | Watering reminders |
| scheduler.go | Watering schedule |
| history.go | Care history |
| report.go | Console report |
| sample.go | Demo plants |

---

### Run

```bash
go run .
```

Everything runs using in-memory sample data.
