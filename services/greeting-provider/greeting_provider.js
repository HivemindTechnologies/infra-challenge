const http = require('http');

const greetings = [
  "Good Day",
  "Buen Día",
  "Bonjour",
  "Buongiorno",
  "Bom Dia",
  "Goede Dag",
  "God Dag",
  "Guten Tag",
];

function getRandomGreeting() {
  const randomIndex = Math.floor(Math.random() * greetings.length);
  return greetings[randomIndex];
}

const server = http.createServer((req, res) => {
  if (req.method === 'GET' && req.url === '/') {
    res.writeHead(200, { 'Content-Type': 'text/plain; charset=utf-8' });
    const greeting = getRandomGreeting();
    console.log(greeting);
    res.end(greeting);
  } else {
    res.writeHead(404, { 'Content-Type': 'text/plain' });
    res.end('404 Not Found');
  }
});

console.log("Hivemind's Greeting Provider")
server.listen(8000)
