use chrono::{Date, Datelike, Utc};
use icu_calendar::DateDurationUnit::{Days, Years};
use icu_calendar::{DateDuration, Iso};

pub fn between_dates(date1: Date<Utc>, date2: Date<Utc>) -> DateDuration<Iso> {
    let iso_date1 = icu_calendar::Date::new_iso_date_from_integers(date1.year(),
                                                     date1.month() as u8,
                                                     date1.day() as u8).unwrap();
    let iso_date2 = icu_calendar::Date::new_iso_date_from_integers(date2.year(),
                                                     date2.month() as u8,
                                                     date2.day() as u8).unwrap();
    iso_date2.until(&iso_date1, Years, Days)
}