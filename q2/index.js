const readline = require('readline');


async function readInput() {
  return new Promise(resolve => {
    let totalBookings = null;
    let bookings = [];
  
    const inputReader = readline.createInterface({
      input: process.stdin,
      output: process.stdout,
    });
  
    inputReader.on('line', line => {
      if (!totalBookings) {
        totalBookings = Number.parseInt(line);
      } else {
        const [checkInDateString, checkOutDateString] = line.split(' ');
        bookings.push({
          checkInDate: new Date(checkInDateString),
          checkOutDate: new Date(checkOutDateString),
        });
        if (bookings.length >= totalBookings) {
          inputReader.close();
          resolve({
            totalBookings,
            bookings
          });
        }
      }
    });
  });
}

function sortBookingsByCheckInDate(bookings) {
  bookings.sort((a, b) => {
    return a.checkInDate.getTime() < b.checkInDate.getTime() ? -1 : 1;
  });
}

function sortBookingsByCheckOutDate(bookings) {
  bookings.sort((a, b) => {
    return a.checkOutDate.getTime() < b.checkOutDate.getTime() ? -1 : 1;
  });
}

function sortEventsByDateAndOccupancyDiff(bookings) {
  bookings.sort((a, b) => {
    if (a.date.getTime() < b.date.getTime()) {
      return -1;
    }
    if (a.date.getTime() > b.date.getTime()) {
      return 1;
    }
    return a.occupancyDiff < b.occupancyDiff ? -1 : 1;
  });
}

function makeEventsFromBookings(bookings) {
  let events = [];

  sortBookingsByCheckInDate(bookings);
  events = events.concat(
    bookings.map(booking => ({ date: booking.checkInDate, occupancyDiff: 1 }))
  );
  
  sortBookingsByCheckOutDate(bookings);
  events = events.concat(
    bookings.map(booking => ({ date: booking.checkOutDate, occupancyDiff: -1 }))
  );

  sortEventsByDateAndOccupancyDiff(events);
  return events;
}

function findMaxOccupancyDateFromEvents(events) {
  let currentOccupancy = 0;
  let maxOccupancy = 0;
  let maxOccupancyDate = null;
  events.forEach(event => {
    currentOccupancy += event.occupancyDiff;
    if (maxOccupancy < currentOccupancy) {
      maxOccupancy = currentOccupancy;
      maxOccupancyDate = event.date;
    }
  });

  return maxOccupancyDate;
}

function findMaxOccupancyDate(bookings) {
  return findMaxOccupancyDateFromEvents(
    makeEventsFromBookings(bookings)
  );
}

function getDateString(dateInstance) {
  return dateInstance.toISOString().split('T')[0];
}

async function main() {
  const {bookings} = await readInput();
  const maxOccupancyDate = findMaxOccupancyDate(bookings);
  console.log(getDateString(maxOccupancyDate));
}

main();
