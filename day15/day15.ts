const parseInput = (rawInput: string) =>
  rawInput
    .split("\n")
    .map((line) => {
      const sensorX = line.match(/Sensor at x=(-?\d+)/)![1];
      const sensorY = line.match(/, y=(-?\d+): /)![1];
      const beaconX = line.match(/beacon is at x=(-?\d+)/)![1];
      const beaconY = line.match(/, y=(-?\d+)$/)![1];
      return {
        sensor: { x: +sensorX, y: +sensorY },
        beacon: { x: +beaconX, y: +beaconY },
      };
    });

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const targetRow = input.length < 30 ? 10 : 2000000;
  const map: Map<number, Map<number, "S" | "B" | "#">> = new Map();

  for (const { sensor, beacon } of input) {
    const distance = Math.abs(beacon.x - sensor.x) +
      Math.abs(beacon.y - sensor.y);
    for (let x = sensor.x - distance; x <= sensor.x + distance; x++) {
      const y = targetRow;
      const thisDistance = Math.abs(sensor.x - x) + Math.abs(sensor.y - y);
      if (thisDistance > distance) {
        continue;
      }
      if (!map.has(y)) {
        map.set(y, new Map());
      }
      map.get(y)!.set(x, map.get(y)!.get(x) || "#");
    }
    if (!map.has(sensor.y)) {
      map.set(sensor.y, new Map());
    }
    if (!map.has(beacon.y)) {
      map.set(beacon.y, new Map());
    }
    map.get(sensor.y)!.set(sensor.x, "S");
    map.get(beacon.y)!.set(beacon.x, "B");
  }

  let freeSpaces = 0;
  for (const [, value] of map.get(targetRow)!) {
    if (value === "#") freeSpaces++;
  }

  console.log("Part 1:", freeSpaces);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const maxCoord = input.length < 30 ? 20 : 4000000;

  const sensors = input.map(({ sensor, beacon }) => {
    return {
      sensor,
      range: Math.abs(beacon.x - sensor.x) + Math.abs(beacon.y - sensor.y),
    };
  });

  for (const { sensor, range } of sensors) {
    const check = { x: sensor.x, y: sensor.y - range - 1 };
    const dir = { x: 1, y: 1 };
    while (true) {
      if (check.x > sensor.x && check.y === sensor.y) {
        dir.x = -1;
      } else if (check.x === sensor.x && check.y > sensor.y) {
        dir.y = -1;
      } else if (check.x < sensor.x && check.y === sensor.y) {
        dir.x = 1;
      } else if (
        dir.y === -1 &&
        check.x === sensor.x &&
        check.y === sensor.y - range - 1
      ) {
        break;
      }
      if (
        !(
          check.x < 0 ||
          check.y < 0 ||
          check.x > maxCoord ||
          check.y > maxCoord
        )
      ) {
        let valid = true;
        for (const otherSensor of sensors) {
          if (otherSensor.sensor === sensor) {
            continue;
          }
          const distance = Math.abs(check.x - otherSensor.sensor.x) +
            Math.abs(check.y - otherSensor.sensor.y);
          if (distance <= otherSensor.range) {
            valid = false;
            break;
          }
        }
        if (valid) {
          const result = (check.x * 4000000 + check.y).toString();
          console.log("Part 2:", result);
          return;
        }
      }
      check.x += dir.x;
      check.y += dir.y;
    }
  }
};

export default function (rawInput: string) {
  part1(rawInput);
  part2(rawInput);
}
