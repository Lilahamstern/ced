<template>
    <v-card>
        <v-card-title>
            Projects
            <v-spacer/>
            <v-text-field v-model="search" append-icon="search" label="Search" singel-line hide-details/>
        </v-card-title>
        <v-data-table
        :headers="headers"
        :items="projects"
        :search="search"
        :single-expand="true"
        show-expand
        :items-per-page='15'
        class="elevation-5"
        >

            <template v-slot:expanded-item="{ headers, item }">
                <template v-for="sub in item.versions" >
                    <td :key="sub.id"></td>
                    <td :key="sub.id">{{ item.version }} </td>
                    <td :key="sub.id">{{ item.name }} </td>
                    <td  :key="sub.id">{{ item.client }} </td>
                    <td :key="sub.id">{{ item.manager }} </td>
                    <td :key="sub.id">{{ item.sector }} </td>
                    <td :key="sub.id">{{ item.co2 }} </td>
                    <td :key="sub.id">{{ item.id}} </td>
                </template>
            </template>
        </v-data-table>
    </v-card>
</template>

<script>
    export default {
        name: "ProjectList",
        data: () => {
            return {
                search: '',
                headers: [
                    {
                        text: 'Version',
                        align: 'start',
                        value: 'version'
                    },
                    {
                        text: 'Name',
                        value: 'name'
                    },
                    {
                        text: 'Client',
                        value: 'client'
                    },
                    {
                        text: 'Manager',
                        value: 'manager'
                    },
                    {
                        text: 'Sector',
                        value: 'sector'
                    },
                    {
                        text: 'Co2/m2',
                        value: 'co2'
                    },
                    {
                        text: 'Database ID',
                        value: 'id'
                    }

            ],
                projects: [
                    {
                        "version": 2,
                        "id": "5eff84446f7d9dd26e47f8e6",
                        "name": "Fenimore Street Penny",
                        "client": "Dadabase",
                        "manager": "David Walls",
                        "sector": "Hillel Place",
                        "co2": -43.87228,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff84447d70bd2baf40e9fb",
                                "name": "Garland Court Elena",
                                "client": "Yogasm",
                                "manager": "Pittman Warren",
                                "sector": "Hutchinson Court",
                                "co2": -82.247525
                            },
                            {
                                "version": 1,
                                "id": "5eff84440a314dffffb33cbb",
                                "name": "Dunham Place Dillard",
                                "client": "Xylar",
                                "manager": "Holden Riley",
                                "sector": "Billings Place",
                                "co2": 131.952656
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff84442cb43b9bab2edcb4",
                        "name": "Loring Avenue Woods",
                        "client": "Overplex",
                        "manager": "Price Olsen",
                        "sector": "Autumn Avenue",
                        "co2": 72.073997,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84447b95b86e737615f5",
                                "name": "Bassett Avenue Atkins",
                                "client": "Netagy",
                                "manager": "Ellis Mckee",
                                "sector": "Danforth Street",
                                "co2": 5.527403
                            },
                            {
                                "version": 3,
                                "id": "5eff844471ac2e80bda5b27f",
                                "name": "Bogart Street Rena",
                                "client": "Goko",
                                "manager": "Ina Hughes",
                                "sector": "Tennis Court",
                                "co2": -117.992983
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444b1c45184d2d97724",
                        "name": "Atlantic Avenue Boyer",
                        "client": "Homelux",
                        "manager": "Gibson Barber",
                        "sector": "Brevoort Place",
                        "co2": 161.810612,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84441926594c78470c84",
                                "name": "Thatford Avenue Maxwell",
                                "client": "Ohmnet",
                                "manager": "Shana Merritt",
                                "sector": "Bulwer Place",
                                "co2": -27.146074
                            },
                            {
                                "version": 3,
                                "id": "5eff844486d3065c580daab5",
                                "name": "Applegate Court Matilda",
                                "client": "Squish",
                                "manager": "Clara Simpson",
                                "sector": "Lorraine Street",
                                "co2": -177.198956
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444d66d0571636b965f",
                        "name": "Paerdegat Avenue Lois",
                        "client": "Geoforma",
                        "manager": "Brandie Gonzalez",
                        "sector": "Downing Street",
                        "co2": 5.053016,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff84440a8f11858996e95c",
                                "name": "Livingston Street Larsen",
                                "client": "Aquazure",
                                "manager": "Dickson Brock",
                                "sector": "Berry Street",
                                "co2": 154.025909
                            },
                            {
                                "version": 4,
                                "id": "5eff84448b1e88b07e8c5bef",
                                "name": "Hampton Place Ila",
                                "client": "Viasia",
                                "manager": "Margaret Carney",
                                "sector": "Colonial Road",
                                "co2": 26.540849
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444fb0f40a61198110a",
                        "name": "Temple Court Kristy",
                        "client": "Grupoli",
                        "manager": "Anderson Graves",
                        "sector": "Forest Place",
                        "co2": 116.335694,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff84446552a7a98613c20c",
                                "name": "Portland Avenue Marcie",
                                "client": "Sentia",
                                "manager": "Rosalie Mclaughlin",
                                "sector": "Kent Avenue",
                                "co2": -155.394703
                            },
                            {
                                "version": 2,
                                "id": "5eff8444cd50e3aeb7156b6e",
                                "name": "Sapphire Street Ericka",
                                "client": "Joviold",
                                "manager": "Velez Sullivan",
                                "sector": "Vanderbilt Avenue",
                                "co2": 158.284426
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff8444e1200e45f560f0e5",
                        "name": "Jodie Court Sara",
                        "client": "Permadyne",
                        "manager": "Hooper Bradford",
                        "sector": "Folsom Place",
                        "co2": 152.678769,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff844457957dbf9eb54248",
                                "name": "Highland Boulevard Norman",
                                "client": "Tubesys",
                                "manager": "Wolfe Daniels",
                                "sector": "Louis Place",
                                "co2": -10.088496
                            },
                            {
                                "version": 1,
                                "id": "5eff8444b5de0af091fd423c",
                                "name": "Wilson Street Rosario",
                                "client": "Idealis",
                                "manager": "Beatriz Eaton",
                                "sector": "Miller Avenue",
                                "co2": -96.722919
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff8444ca1d3b6264d3be00",
                        "name": "Division Avenue Benjamin",
                        "client": "Magmina",
                        "manager": "Jo Sims",
                        "sector": "Hemlock Street",
                        "co2": -37.326284,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84440ed6802439200708",
                                "name": "Atkins Avenue Clarissa",
                                "client": "Qiao",
                                "manager": "Rosalyn Vargas",
                                "sector": "Moffat Street",
                                "co2": -84.363156
                            },
                            {
                                "version": 4,
                                "id": "5eff844477eacee6f70321fa",
                                "name": "Montrose Avenue Durham",
                                "client": "Eclipto",
                                "manager": "Ayers Lucas",
                                "sector": "Monaco Place",
                                "co2": 109.676246
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444044237198eaa85d0",
                        "name": "Interborough Parkway Austin",
                        "client": "Kiggle",
                        "manager": "Lena Ross",
                        "sector": "Denton Place",
                        "co2": -134.636268,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff8444e31d89ac57d1f3f4",
                                "name": "Chase Court Katharine",
                                "client": "Eplosion",
                                "manager": "Frost Hartman",
                                "sector": "Box Street",
                                "co2": -93.451389
                            },
                            {
                                "version": 4,
                                "id": "5eff8444b7f0533dc48128b2",
                                "name": "Hart Place Claire",
                                "client": "Emtrak",
                                "manager": "Kramer Jacobs",
                                "sector": "Times Placez",
                                "co2": 43.965828
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff8444d32a189ca05503c0",
                        "name": "Greenwood Avenue House",
                        "client": "Rockabye",
                        "manager": "Marina Fry",
                        "sector": "Perry Place",
                        "co2": -50.716757,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444a5e39ce88497c3f6",
                                "name": "Nevins Street Haley",
                                "client": "Zolarex",
                                "manager": "Lewis Reyes",
                                "sector": "Bartlett Place",
                                "co2": 56.930802
                            },
                            {
                                "version": 2,
                                "id": "5eff844433bff0a72641cf3e",
                                "name": "Calyer Street Hinton",
                                "client": "Limozen",
                                "manager": "Blake Meyer",
                                "sector": "Montauk Court",
                                "co2": -132.909276
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff84445d53431a7d2d9bdf",
                        "name": "Hinckley Place Winters",
                        "client": "Entality",
                        "manager": "Gena Richard",
                        "sector": "Jackson Place",
                        "co2": -119.810851,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff8444ec54147670ac7da3",
                                "name": "Garfield Place Melba",
                                "client": "Suremax",
                                "manager": "Aguirre Roman",
                                "sector": "Bancroft Place",
                                "co2": 93.074811
                            },
                            {
                                "version": 4,
                                "id": "5eff8444f42a8031ff420b44",
                                "name": "Troutman Street Brooks",
                                "client": "Isoternia",
                                "manager": "Weber Mcdonald",
                                "sector": "Duffield Street",
                                "co2": -172.231593
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff84447fc9b0bb2731001a",
                        "name": "Seba Avenue Greta",
                        "client": "Endipine",
                        "manager": "Camacho Suarez",
                        "sector": "Sutter Avenue",
                        "co2": -134.484719,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444b13c3703551016af",
                                "name": "Newkirk Placez Savage",
                                "client": "Geologix",
                                "manager": "Hope Jones",
                                "sector": "Nixon Court",
                                "co2": -47.665741
                            },
                            {
                                "version": 2,
                                "id": "5eff8444180ef803f72f6dcc",
                                "name": "Lake Avenue Bonnie",
                                "client": "Xeronk",
                                "manager": "Landry Mcgee",
                                "sector": "Quentin Road",
                                "co2": -99.189723
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444235c6ec472cd1386",
                        "name": "Taaffe Place Porter",
                        "client": "Geeketron",
                        "manager": "Macias Howe",
                        "sector": "Anthony Street",
                        "co2": -78.037126,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff84446f5c946a2bbd089b",
                                "name": "Chauncey Street Duran",
                                "client": "Uni",
                                "manager": "Patton Guzman",
                                "sector": "Judge Street",
                                "co2": 23.550947
                            },
                            {
                                "version": 4,
                                "id": "5eff844432a7d2713f16bf4c",
                                "name": "Whitwell Place William",
                                "client": "Roboid",
                                "manager": "Brittney Matthews",
                                "sector": "Apollo Street",
                                "co2": 33.343691
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff84449d363d444b025cf1",
                        "name": "Halleck Street Gwendolyn",
                        "client": "Multron",
                        "manager": "Hodge Black",
                        "sector": "Bergen Street",
                        "co2": 43.107314,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff8444d0fe80d97fec58ab",
                                "name": "Cropsey Avenue Diana",
                                "client": "Genekom",
                                "manager": "Natalie Moore",
                                "sector": "Rost Place",
                                "co2": -123.528984
                            },
                            {
                                "version": 2,
                                "id": "5eff84447618a817ba7dcbba",
                                "name": "Gain Court Franklin",
                                "client": "Zilphur",
                                "manager": "Hill Beach",
                                "sector": "Coventry Road",
                                "co2": 24.857592
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444e55b8e5c64a6dec6",
                        "name": "Hazel Court Stephanie",
                        "client": "Maxemia",
                        "manager": "Alta Cotton",
                        "sector": "Fuller Place",
                        "co2": 17.246213,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff844489f7c9a93954a55e",
                                "name": "Plaza Street Salas",
                                "client": "Vixo",
                                "manager": "Sawyer Whitney",
                                "sector": "Cranberry Street",
                                "co2": 125.957475
                            },
                            {
                                "version": 3,
                                "id": "5eff8444252e019edb54f15a",
                                "name": "Brooklyn Road Travis",
                                "client": "Grok",
                                "manager": "Lori Beard",
                                "sector": "Durland Place",
                                "co2": -21.756069
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444a8ec20ecca91d168",
                        "name": "Middleton Street Rosemary",
                        "client": "Pigzart",
                        "manager": "Jean Weiss",
                        "sector": "Cox Place",
                        "co2": 57.278892,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444e2a7db084a178c38",
                                "name": "Sackett Street Cathryn",
                                "client": "Solgan",
                                "manager": "Jaime Cortez",
                                "sector": "Madeline Court",
                                "co2": -14.895987
                            },
                            {
                                "version": 2,
                                "id": "5eff84447ce7e810f18f788d",
                                "name": "Clara Street Marie",
                                "client": "Equicom",
                                "manager": "Renee Giles",
                                "sector": "Bergen Place",
                                "co2": 115.127385
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444ba4e3b8386eae8f7",
                        "name": "Ovington Court Moss",
                        "client": "Datacator",
                        "manager": "Gloria Everett",
                        "sector": "Jackson Street",
                        "co2": -67.87879,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff844460123172984f8a70",
                                "name": "Provost Street Dorthy",
                                "client": "Proxsoft",
                                "manager": "Rush Price",
                                "sector": "Imlay Street",
                                "co2": -76.537089
                            },
                            {
                                "version": 2,
                                "id": "5eff8444475b007b8c6a3abb",
                                "name": "Buffalo Avenue Nguyen",
                                "client": "Cinaster",
                                "manager": "Delia Patrick",
                                "sector": "Ocean Avenue",
                                "co2": -99.755742
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff84449ab8759a5c52b378",
                        "name": "Arkansas Drive Alyssa",
                        "client": "Webiotic",
                        "manager": "Aisha Goodwin",
                        "sector": "Garden Street",
                        "co2": -146.457572,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444d425d7d201f031d7",
                                "name": "Clay Street Jefferson",
                                "client": "Automon",
                                "manager": "Haley Barrera",
                                "sector": "Nelson Street",
                                "co2": -34.596163
                            },
                            {
                                "version": 1,
                                "id": "5eff844489f1679d2dc6a599",
                                "name": "Baycliff Terrace Zelma",
                                "client": "Candecor",
                                "manager": "Lawson Guthrie",
                                "sector": "Royce Street",
                                "co2": 106.221628
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff844457a05accf3d7de1e",
                        "name": "Regent Place Ballard",
                        "client": "Poshome",
                        "manager": "Neva Hardin",
                        "sector": "Montgomery Place",
                        "co2": -69.250525,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84448fb54348e3cd97d3",
                                "name": "Juliana Place Donovan",
                                "client": "Retrotex",
                                "manager": "Oneill Cohen",
                                "sector": "Degraw Street",
                                "co2": -150.004002
                            },
                            {
                                "version": 2,
                                "id": "5eff8444efc84863d064f8f8",
                                "name": "Furman Avenue Avila",
                                "client": "Anivet",
                                "manager": "Jackson Snow",
                                "sector": "Bedford Place",
                                "co2": 145.964855
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff84445d6d73705a036176",
                        "name": "Herzl Street Murray",
                        "client": "Maroptic",
                        "manager": "Morgan Harding",
                        "sector": "Court Square",
                        "co2": 140.797746,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff8444fe91b1877181a425",
                                "name": "Charles Place Leila",
                                "client": "Shadease",
                                "manager": "Cynthia Cherry",
                                "sector": "Schweikerts Walk",
                                "co2": -152.757083
                            },
                            {
                                "version": 1,
                                "id": "5eff84449e0a6e0351b1b366",
                                "name": "Lawrence Street Whitfield",
                                "client": "Zenolux",
                                "manager": "Sweet Brennan",
                                "sector": "Poplar Avenue",
                                "co2": 127.947537
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff8444de97f3142538b60a",
                        "name": "Church Lane Mcguire",
                        "client": "Tellifly",
                        "manager": "Adele Marsh",
                        "sector": "Euclid Avenue",
                        "co2": 27.600761,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff8444bd91fef7e733b8dd",
                                "name": "Heath Place Margie",
                                "client": "Quizmo",
                                "manager": "Tracie Bird",
                                "sector": "Montana Place",
                                "co2": 73.922743
                            },
                            {
                                "version": 4,
                                "id": "5eff8444a9b04e1c29f8fa07",
                                "name": "Commercial Street Beth",
                                "client": "Otherway",
                                "manager": "Kristina Schroeder",
                                "sector": "Stockholm Street",
                                "co2": 26.036646
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff844482a99ac3e00170a6",
                        "name": "Cadman Plaza Kasey",
                        "client": "Junipoor",
                        "manager": "Horne Freeman",
                        "sector": "Columbia Street",
                        "co2": 166.332622,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84440aa21e97c24eebe1",
                                "name": "Tech Place Lucille",
                                "client": "Sultraxin",
                                "manager": "Callahan Horton",
                                "sector": "Dennett Place",
                                "co2": -42.720486
                            },
                            {
                                "version": 4,
                                "id": "5eff84444cfbe634b03146d8",
                                "name": "Quay Street Michelle",
                                "client": "Panzent",
                                "manager": "Murphy Clements",
                                "sector": "Dahill Road",
                                "co2": -151.436972
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff844471fa876069ebf2f2",
                        "name": "Harrison Avenue Jill",
                        "client": "Bitrex",
                        "manager": "Mccoy Powell",
                        "sector": "Branton Street",
                        "co2": -103.993587,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff8444568d524103bf24b3",
                                "name": "Grafton Street Fry",
                                "client": "Digique",
                                "manager": "Franco Murphy",
                                "sector": "George Street",
                                "co2": 116.10596
                            },
                            {
                                "version": 1,
                                "id": "5eff8444c345b3e4f0681942",
                                "name": "Quentin Street Dolly",
                                "client": "Ontagene",
                                "manager": "Park Beasley",
                                "sector": "Radde Place",
                                "co2": -139.681819
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444c1b738ba46a714a6",
                        "name": "Roosevelt Place Ofelia",
                        "client": "Corepan",
                        "manager": "Wiley Rosa",
                        "sector": "Navy Walk",
                        "co2": -50.80559,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff8444631caf4da13efe01",
                                "name": "Alabama Avenue Whitney",
                                "client": "Rooforia",
                                "manager": "Rice Johns",
                                "sector": "Richards Street",
                                "co2": -33.827362
                            },
                            {
                                "version": 2,
                                "id": "5eff84446f21eb5f6fb207a0",
                                "name": "Williams Place Madelyn",
                                "client": "Andershun",
                                "manager": "Michael Knight",
                                "sector": "Meserole Street",
                                "co2": -60.003336
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff84441bf4bab562e6eafb",
                        "name": "Duryea Court Mathews",
                        "client": "Crustatia",
                        "manager": "Ashley Holder",
                        "sector": "Boulevard Court",
                        "co2": -123.768473,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84445d011989d6d63d9c",
                                "name": "Dupont Street Lisa",
                                "client": "Chorizon",
                                "manager": "Maggie Vance",
                                "sector": "Norfolk Street",
                                "co2": -87.042995
                            },
                            {
                                "version": 1,
                                "id": "5eff8444f5ad38d9e99bb8c1",
                                "name": "Caton Place Loraine",
                                "client": "Corecom",
                                "manager": "Graves Tyson",
                                "sector": "Freeman Street",
                                "co2": 163.417989
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff844410fad28d14898511",
                        "name": "Maple Street Maddox",
                        "client": "Navir",
                        "manager": "Mabel Lester",
                        "sector": "Newel Street",
                        "co2": -56.985155,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff844432ee8e2f092957ff",
                                "name": "Richardson Street Melendez",
                                "client": "Interfind",
                                "manager": "Monica Baxter",
                                "sector": "Leonard Street",
                                "co2": -59.70341
                            },
                            {
                                "version": 2,
                                "id": "5eff8444b364a84ed82cd997",
                                "name": "Luquer Street Blanche",
                                "client": "Miracula",
                                "manager": "Owens Strong",
                                "sector": "Kenilworth Place",
                                "co2": -20.636758
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff8444df80a92a547dc03a",
                        "name": "Newkirk Avenue Gardner",
                        "client": "Signity",
                        "manager": "Golden Marquez",
                        "sector": "Bush Street",
                        "co2": -166.723614,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444abfad7af519f9483",
                                "name": "Brooklyn Avenue Victoria",
                                "client": "Acumentor",
                                "manager": "Wolf Barr",
                                "sector": "Irving Street",
                                "co2": -149.344461
                            },
                            {
                                "version": 1,
                                "id": "5eff844423f68a78ea423710",
                                "name": "Louisa Street Knight",
                                "client": "Firewax",
                                "manager": "Dudley Abbott",
                                "sector": "Gates Avenue",
                                "co2": 144.184639
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff844496201e8e2a5035f9",
                        "name": "Sutton Street Swanson",
                        "client": "Zerology",
                        "manager": "Benita Joyner",
                        "sector": "Monroe Street",
                        "co2": -145.921325,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff8444a0aaa4f0de5eb8f2",
                                "name": "Sands Street Cortez",
                                "client": "Combogene",
                                "manager": "Spencer Zimmerman",
                                "sector": "Beverley Road",
                                "co2": -11.247062
                            },
                            {
                                "version": 3,
                                "id": "5eff84444238f6093cdbfd1b",
                                "name": "Bank Street Allyson",
                                "client": "Hyplex",
                                "manager": "Sweeney Whitehead",
                                "sector": "Vandervoort Avenue",
                                "co2": 124.080971
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff8444983b3ca38d31b078",
                        "name": "Channel Avenue Juliana",
                        "client": "Pyramia",
                        "manager": "Ines Flores",
                        "sector": "Empire Boulevard",
                        "co2": -119.875634,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff8444a2c169d1eb5c4e88",
                                "name": "Melba Court Kelley",
                                "client": "Zentry",
                                "manager": "Allison Perez",
                                "sector": "Remsen Street",
                                "co2": 108.761408
                            },
                            {
                                "version": 2,
                                "id": "5eff8444bdf719c656c2f03f",
                                "name": "Malta Street Leola",
                                "client": "Amril",
                                "manager": "Kerry Huff",
                                "sector": "Reeve Place",
                                "co2": 34.315562
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444137e31c03adb9825",
                        "name": "Boynton Place Nina",
                        "client": "Steelfab",
                        "manager": "Willa Britt",
                        "sector": "Pleasant Place",
                        "co2": 104.529328,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84448c95f35428a315ea",
                                "name": "Sedgwick Place Gwen",
                                "client": "Globoil",
                                "manager": "Alejandra Soto",
                                "sector": "Crawford Avenue",
                                "co2": -3.263525
                            },
                            {
                                "version": 4,
                                "id": "5eff8444066674ae6430fb2f",
                                "name": "Gardner Avenue Carmella",
                                "client": "Imkan",
                                "manager": "Bobbie Booth",
                                "sector": "Tompkins Place",
                                "co2": 159.105876
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff84447a7a174120cfb93b",
                        "name": "Alton Place Conner",
                        "client": "Pawnagra",
                        "manager": "Carr Frazier",
                        "sector": "Williamsburg Street",
                        "co2": -117.084545,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84445aedbd85632d2038",
                                "name": "Havemeyer Street Phoebe",
                                "client": "Songlines",
                                "manager": "Berger Love",
                                "sector": "Ridgewood Place",
                                "co2": 95.872493
                            },
                            {
                                "version": 4,
                                "id": "5eff8444a8d4ff241e32ee57",
                                "name": "Voorhies Avenue Annmarie",
                                "client": "Accuprint",
                                "manager": "Tara Rodgers",
                                "sector": "Lois Avenue",
                                "co2": 102.608859
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff844476849fbfd982c25d",
                        "name": "Havens Place Luna",
                        "client": "Eventage",
                        "manager": "Stanton Peters",
                        "sector": "Montieth Street",
                        "co2": -63.734334,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84446a19ae03637f446c",
                                "name": "Turner Place Waters",
                                "client": "Magnemo",
                                "manager": "Paulette Crosby",
                                "sector": "Borinquen Pl",
                                "co2": 96.129312
                            },
                            {
                                "version": 4,
                                "id": "5eff84444e5bca8c8843a1ba",
                                "name": "Senator Street Herminia",
                                "client": "Opticom",
                                "manager": "Ford Petersen",
                                "sector": "Vandervoort Place",
                                "co2": -108.679354
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff84447aed9bc33fdcf3dc",
                        "name": "Huntington Street Araceli",
                        "client": "Pearlessa",
                        "manager": "Long Howell",
                        "sector": "Glen Street",
                        "co2": 63.375803,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff844470dbd19e5ee80737",
                                "name": "Rockwell Place Snyder",
                                "client": "Zentury",
                                "manager": "Georgina Dominguez",
                                "sector": "Flatbush Avenue",
                                "co2": -79.598541
                            },
                            {
                                "version": 1,
                                "id": "5eff84441d2d7511f331d3bc",
                                "name": "Carroll Street Pauline",
                                "client": "Perkle",
                                "manager": "Quinn Morin",
                                "sector": "Rewe Street",
                                "co2": 49.215981
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444a80af97393f0880e",
                        "name": "Pioneer Street Cohen",
                        "client": "Sportan",
                        "manager": "Sykes Robinson",
                        "sector": "Miami Court",
                        "co2": -4.765902,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff8444a5779a948b0a23d0",
                                "name": "Vandalia Avenue Myers",
                                "client": "Deepends",
                                "manager": "Burnett Griffin",
                                "sector": "Aster Court",
                                "co2": 29.059763
                            },
                            {
                                "version": 3,
                                "id": "5eff8444cf9b8005f5993718",
                                "name": "Windsor Place Addie",
                                "client": "Zipak",
                                "manager": "Bertie Vaughn",
                                "sector": "Banker Street",
                                "co2": 2.275265
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff844434619a85aa4aca40",
                        "name": "Java Street Nielsen",
                        "client": "Vantage",
                        "manager": "Leticia Bauer",
                        "sector": "Meserole Avenue",
                        "co2": -94.42144,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84443e3a39e978ad5608",
                                "name": "Lafayette Avenue Krystal",
                                "client": "Kidstock",
                                "manager": "Adrienne Bean",
                                "sector": "Victor Road",
                                "co2": 134.807521
                            },
                            {
                                "version": 3,
                                "id": "5eff8444328830f223a2c615",
                                "name": "Duryea Place Vang",
                                "client": "Biflex",
                                "manager": "Dana Stark",
                                "sector": "Madoc Avenue",
                                "co2": 4.900764
                            }
                        ]
                    },
                    {
                        "version": 2,
                        "id": "5eff8444587a5c572f7a2850",
                        "name": "Vanderveer Place Tyson",
                        "client": "Isonus",
                        "manager": "Maribel Wilkerson",
                        "sector": "Farragut Road",
                        "co2": 26.456459,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff8444346b88ee83805026",
                                "name": "Hegeman Avenue Rachelle",
                                "client": "Orbaxter",
                                "manager": "Faye Juarez",
                                "sector": "Bragg Court",
                                "co2": -9.745039
                            },
                            {
                                "version": 4,
                                "id": "5eff84443575690658c1c116",
                                "name": "Lloyd Street Connie",
                                "client": "Venoflex",
                                "manager": "Mccormick Cole",
                                "sector": "Middagh Street",
                                "co2": 117.960257
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff84447767e6044acf1f39",
                        "name": "Mill Road Bowers",
                        "client": "Gronk",
                        "manager": "Bridgett Bryant",
                        "sector": "Metrotech Courtr",
                        "co2": 158.642554,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff8444b87287899e5dc4e9",
                                "name": "Hamilton Avenue Gabrielle",
                                "client": "Uxmox",
                                "manager": "Nikki Leblanc",
                                "sector": "Ferris Street",
                                "co2": -145.092514
                            },
                            {
                                "version": 1,
                                "id": "5eff8444698f5d6a73308127",
                                "name": "Richmond Street Sargent",
                                "client": "Gushkool",
                                "manager": "Annabelle Simon",
                                "sector": "Ocean Court",
                                "co2": 126.216837
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff84448e4c8729fc772ee5",
                        "name": "Aberdeen Street Loretta",
                        "client": "Undertap",
                        "manager": "Mccarty Hicks",
                        "sector": "Montague Terrace",
                        "co2": 123.676017,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff84446b789ea558a92557",
                                "name": "Chester Street Abigail",
                                "client": "Earthpure",
                                "manager": "Jerri Watkins",
                                "sector": "Maujer Street",
                                "co2": -84.88842
                            },
                            {
                                "version": 1,
                                "id": "5eff84441372f19dea84b2b0",
                                "name": "Woodruff Avenue Raymond",
                                "client": "Emtrac",
                                "manager": "Carol Beck",
                                "sector": "Varick Avenue",
                                "co2": -65.958618
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff8444d9a97c90f8a323ea",
                        "name": "Doone Court Ruiz",
                        "client": "Incubus",
                        "manager": "Angel Robertson",
                        "sector": "Leonora Court",
                        "co2": 166.821217,
                        "versions": [
                            {
                                "version": 1,
                                "id": "5eff84445bb0231722b25c62",
                                "name": "Rapelye Street Campos",
                                "client": "Boilicon",
                                "manager": "Lavonne Garcia",
                                "sector": "Fayette Street",
                                "co2": -63.45379
                            },
                            {
                                "version": 1,
                                "id": "5eff84446c3822f3b015cbaa",
                                "name": "Hendrix Street Becker",
                                "client": "Accel",
                                "manager": "Graham Hoover",
                                "sector": "Clifton Place",
                                "co2": 10.310011
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444907dd56bc4059729",
                        "name": "Visitation Place Henson",
                        "client": "Kaggle",
                        "manager": "Effie Pate",
                        "sector": "Caton Avenue",
                        "co2": -103.277589,
                        "versions": [
                            {
                                "version": 3,
                                "id": "5eff8444ba09cc33a0b0f231",
                                "name": "Lacon Court Mcmahon",
                                "client": "Zensure",
                                "manager": "Isabel Burton",
                                "sector": "Bartlett Street",
                                "co2": -162.282678
                            },
                            {
                                "version": 3,
                                "id": "5eff84440e820a60aeff9461",
                                "name": "Irwin Street Weeks",
                                "client": "Ginkogene",
                                "manager": "Francis Stein",
                                "sector": "Kimball Street",
                                "co2": -12.823415
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff8444af064d85a7fe6ebb",
                        "name": "Woodhull Street Gamble",
                        "client": "Rockyard",
                        "manager": "Rhea Gallegos",
                        "sector": "Stoddard Place",
                        "co2": 78.586236,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff844407a1bd592e024030",
                                "name": "Hale Avenue Sheree",
                                "client": "Codax",
                                "manager": "Millicent Kirby",
                                "sector": "Crescent Street",
                                "co2": -172.611351
                            },
                            {
                                "version": 1,
                                "id": "5eff8444bb3e467f8c2c1474",
                                "name": "Harden Street Josie",
                                "client": "Printspan",
                                "manager": "Noelle Wallace",
                                "sector": "Will Place",
                                "co2": -112.818638
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff84444364ad3584001fa1",
                        "name": "Forrest Street Leslie",
                        "client": "Isotrack",
                        "manager": "Kelley Ward",
                        "sector": "Lawn Court",
                        "co2": 174.47654,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff84444db9285eb1a4db0f",
                                "name": "Kane Place Pearson",
                                "client": "Xsports",
                                "manager": "Dejesus Clark",
                                "sector": "Harwood Place",
                                "co2": 56.507012
                            },
                            {
                                "version": 3,
                                "id": "5eff8444ed017e0343526eb6",
                                "name": "Engert Avenue Spears",
                                "client": "Nitracyr",
                                "manager": "Mai Norman",
                                "sector": "Corbin Place",
                                "co2": -2.916073
                            }
                        ]
                    },
                    {
                        "version": 4,
                        "id": "5eff8444e15155265531e0f5",
                        "name": "Saratoga Avenue Jensen",
                        "client": "Extrawear",
                        "manager": "Edwina Barnett",
                        "sector": "Oriental Court",
                        "co2": -163.308133,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff8444ff13e350646d2648",
                                "name": "Boerum Place Rios",
                                "client": "Speedbolt",
                                "manager": "Sue Sanders",
                                "sector": "Oxford Street",
                                "co2": 173.000945
                            },
                            {
                                "version": 4,
                                "id": "5eff8444fd8fd2779c0212e8",
                                "name": "Neptune Avenue Mack",
                                "client": "Helixo",
                                "manager": "Holt Brewer",
                                "sector": "Colonial Court",
                                "co2": 126.025909
                            }
                        ]
                    },
                    {
                        "version": 1,
                        "id": "5eff8444ad2ac0088d256143",
                        "name": "Lewis Avenue Lyons",
                        "client": "Wrapture",
                        "manager": "Anna Higgins",
                        "sector": "Conover Street",
                        "co2": -179.994092,
                        "versions": [
                            {
                                "version": 4,
                                "id": "5eff844498e563cf7beba049",
                                "name": "Knight Court Veronica",
                                "client": "Noralex",
                                "manager": "Marcella Mccarthy",
                                "sector": "Franklin Avenue",
                                "co2": 153.161031
                            },
                            {
                                "version": 1,
                                "id": "5eff844445d2d200dd1cf03f",
                                "name": "Christopher Avenue Irwin",
                                "client": "Trollery",
                                "manager": "Lesa Berry",
                                "sector": "Clarendon Road",
                                "co2": -4.717616
                            }
                        ]
                    },
                    {
                        "version": 3,
                        "id": "5eff8444e034a204a3466796",
                        "name": "Chestnut Avenue Tania",
                        "client": "Geeky",
                        "manager": "Velazquez Rodriquez",
                        "sector": "Hoyts Lane",
                        "co2": -147.888443,
                        "versions": [
                            {
                                "version": 2,
                                "id": "5eff8444846deed1716e9497",
                                "name": "Halsey Street Nettie",
                                "client": "Comtext",
                                "manager": "Simon Pennington",
                                "sector": "Malbone Street",
                                "co2": -25.418519
                            },
                            {
                                "version": 1,
                                "id": "5eff8444148f5632b36adb41",
                                "name": "Prospect Place Wall",
                                "client": "Beadzza",
                                "manager": "Griffith Heath",
                                "sector": "Dobbin Street",
                                "co2": 68.564168
                            }
                        ]
                    }
                ]
            }
        }
    }
</script>

<style scoped>

</style>