const firstNames = [
  'John',
  'Paul',
  'George',
  'Ringo',
  'Francis',
  'Jean',
  'Pierre',
  'Marie',
  'Anne',
  'Lucas',
  'Liam',
  'Emma',
  'Olivia',
  'Ava',
  'Isabella',
  'Sophia',
  'Mia',
  'Charlotte',
  'Amelia',
  'Harper',
]

const lastNames = [
  'Lennon',
  'Sykes',
  'Ford',
  'McCartney',
  'Harrison',
  'Starr',
  'Doe',
  'Smith',
  'Johnson',
  'Williams',
  'Brown',
  'Jones',
  'Garcia',
  'Miller',
  'Davis',
  'Rodriguez',
  'Martinez',
  'Peterson',
]

const choice = (arr) => arr[Math.floor(Math.random() * arr.length)]
const cpf = () => {
  let str = ''
  for (let i = 0; i < 11; i++) {
    str += Math.floor(Math.random() * 10)
  }

  return str
}

async function main() {
  for (let i = 0; i < 100; i++) {
    const firstName = choice(firstNames)
    const lastName = choice(lastNames)

    const payload = {
      name: `${firstName} ${lastName}`,
      cpf: cpf(),
    }

    const response = await fetch('http://localhost:8080/v1/accounts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    if (!response.ok) {
      console.error(`Error creating account: ${response.status}`)
      continue
    }

    const account = await response.text()

    console.log(account)
  }
}

main()
