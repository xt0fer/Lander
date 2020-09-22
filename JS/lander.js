var t1 = new Terminal()
var gameAlive = true;

t1.setHeight("600px")
t1.setWidth('800px')


function randomheight() {
    let max = 20000;
    let min = 10000;
    let r = Math.floor(Math.random() * (max - min)) + min;
    return (r % 15000 + 4000)
}

const Gravity = 100 /* The rate in which the spaceship descents in free fall (in ten seconds) */

const version = "1.2"; /* The Version of the program */
const dead = "\nThere were no survivors.\n\n";
const crashed = "\nThe Spaceship crashed. Good luck getting back home.\n\n";
const success = "\nYou made it! Good job!\n\n";
const emptyfuel = "\nThere is no fuel left. You're floating around like Wheatley.\n\n";



const redFormatNumber = "\x1b[31m%d\x1b[0m\t\t";
const greenFormatNumber = "\x1b[32m%d\x1b[0m\t\t";


function gameHeader() {
    s = "";
    s = s + "\nLunar Lander - Version " + version + "\n";
    s = s + "This is a computer simulation of an Apollo lunar landing capsule.\n";
    s = s + "The on-board computer has failed so you have to land the capsule manually.\n";
    s = s + "Set burn rate of retro rockets to any value between 0 (free fall) and 200\n";
    s = s + "(maximum burn) kilo per second. Set burn rate every 10 seconds.\n"; /* That's why we have to go with 10 second-steps. */
    s = s + "You must land at a speed of 2 or 1. Good Luck!\n\n";
    return s;
}

function getHeader() {
    s = "";
    s = s + "\nTime\t";
    s = s + "Speed\t\t";
    s = s + "Fuel\t\t";
    s = s + "Height\t\t";
    s = s + "Burn\n";
    s = s + "----\t";
    s = s + "-----\t\t";
    s = s + "----\t\t";
    s = s + "------\t\t";
    s = s + "----\n";
    return s;
}

// function getBurnRate() int {
//     let burn = 0
//     let junk = ""
//         // Do very simple input validation.
//     for {
//         _,
//         err: = fmt.Scanf("%d", & burn)
//         if err != nil {
//             fmt.Printf("Burn rate needs to be a number.\n>> ")
//             fmt.Scanf("%s", & junk) /* read rest of line if first parse isn't an int.*/
//             continue
//         }
//         if burn < 0 || burn > 200 { /* If there is a wrong entry */
//             fmt.Printf("The burn rate rate must be between 0 and 200.\n> ")
//             continue
//         } else {
//             break
//         }
//     }
//     return burn
// }

function computeDeltaV(vehicle) {
    return (vehicle.Speed + Gravity - vehicle.Burn)
}

function checkStatus(vehicle) {
    s = "";
    if (vehicle.Height <= 0) {
        if (vehicle.Speed > 10) {
            s = dead;
        }
        if (vehicle.Speed < 10 && vehicle.Speed > 3) {
            s = crashed;
        }

        if (vehicle.Speed < 3) {
            s = success;
        }
    } else {
        if (vehicle.Height > 0) {
            s = emptyfuel;
        }
    }
    return s
}

function adjustForBurn(vehicle) {
    console.log("adjustForBurn", vehicle.Burn);
    vehicle.PrevHeight = vehicle.Height;
    vehicle.Speed = computeDeltaV(vehicle);
    vehicle.Height = vehicle.Height - vehicle.Speed;
    vehicle.Fuel = vehicle.Fuel - vehicle.Burn;
}

function stillFlying() {
    return (vehicle.Height > 0);
}

function outOfFuel(vehicle) {
    return (vehicle.Fuel <= 0);
}

function getStatusLine(vehicle) {
    let s = "";
    s = vehicle.Tensec + "0 " + vehicle.Speed + " " + vehicle.Fuel + " " +
        vehicle.Height;
    return s
}

function printRunningHeader(vehicle) {
    if (vehicle.Step >= 24) {
        s = s + "\nTime\t";
        s = s + "Speed\t\t";
        s = s + "Fuel\t\t";
        s = s + "Height\t\t";
        s = s + "Burn\n";
        s = s + "----\t";
        s = s + "-----\t\t";
        s = s + "----\t\t";
        s = s + "------\t\t";
        s = s + "----\n";
        vehicle.Step = 1
    } else {
        if (vehicle.Step < 24) {
            vehicle.Step++
        }
    }
}

function printString(t, string) {
    let a = string.split(/\r?\n/);
    for (i = 0; i < a.length; i++) {
        console.log(">>", a[i]);
        t.print(a[i]);
    }
}

var vehicle = {
    Height: 8000,
    Speed: 1000,
    Fuel: 12000,
    Tensec: 0,
    Burn: 0,
    PrevHeight: 8000,
    Step: 1,
}

// main game loop
async function runGame() {
    let status = ""
    document.body.appendChild(t1.html)

    /* Set initial vehicle parameters */
    let h = randomheight()
    vehicle.Height = h;
    vehicle.PrevHeight = h;

    //t1.print(gameHeader());
    printString(t1, gameHeader());
    //t1.print(getHeader());
    printString(t1, getHeader());
    console.log("tock");

    while (stillFlying() === true) {
        console.log(vehicle);
        // printRunningHeader(vehicle)

        status = getStatusLine(vehicle);
        console.log(status);
        //t1.print(status);
        vehicle.Burn = await t1.input(status, function(input) {
            //t1.print(input);
            return input;
        });

        adjustForBurn(vehicle);

        if (outOfFuel(vehicle) === true) {
            break;
        }
        vehicle.Tensec++;


    }
    status = checkStatus(vehicle);
    t1.print(status);
}