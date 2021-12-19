package controller

const DEMO_RECIPES = `[
  {
    "title":"Julia Child's Coq au Vin",
    "picture":"https://spoonacular.com/recipeImages/648638-556x370.jpg",
    "ingredients":[
      "bacon",
      "bay leaves",
      "butter",
      "button mushrooms",
      "chicken legs",
      "chicken stock",
      "fresh parsley",
      "garlic cloves",
      "parsley",
      "pearl onions",
      "red wine",
      "salt and pepper",
      "thyme sprigs"
    ],
    "ingredientInfo":{
      "bacon":{
        "amount":0.5,
        "measuresUnit":"lb",
        "image":"raw-bacon.png",
        "name":"bacon"
      },
      "bay leaves":{
        "amount":2,
        "measuresUnit":"",
        "image":"bay-leaves.jpg",
        "name":"bay leaves"
      },
      "butter":{
        "amount":2,
        "measuresUnit":"Tbsp",
        "image":"butter-sliced.jpg",
        "name":"butter"
      },
      "button mushrooms":{
        "amount":0.5,
        "measuresUnit":"lb",
        "image":"mushrooms-white.jpg",
        "name":"button mushrooms"
      },
      "chicken legs":{
        "amount":3,
        "measuresUnit":"lbs",
        "image":"chicken-drumsticks.jpg",
        "name":"chicken legs"
      },
      "chicken stock":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"chicken-broth.png",
        "name":"chicken stock"
      },
      "fresh parsley":{
        "amount":6,
        "measuresUnit":"servings",
        "image":"parsley.jpg",
        "name":"fresh parsley"
      },
      "garlic cloves":{
        "amount":6,
        "measuresUnit":"",
        "image":"garlic.png",
        "name":"garlic cloves"
      },
      "parsley":{
        "amount":1,
        "measuresUnit":"sprigs",
        "image":"parsley.jpg",
        "name":"parsley"
      },
      "pearl onions":{
        "amount":20,
        "measuresUnit":"large",
        "image":"pearl-onions.png",
        "name":"pearl onions"
      },
      "red wine":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"red-wine.jpg",
        "name":"red wine"
      },
      "salt and pepper":{
        "amount":6,
        "measuresUnit":"servings",
        "image":"salt-and-pepper.jpg",
        "name":"salt and pepper"
      },
      "thyme sprigs":{
        "amount":1,
        "measuresUnit":"sprigs",
        "image":"thyme.jpg",
        "name":"thyme sprigs"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eBlanch the bacon to remove some of its saltiness. Drop the bacon into a saucepan of cold water, covered by a couple of inches. Bring to a boil, simmer for 5 minutes, drain. Rinse in cold water, pat dry with paper towels. Cut the bacon into 1 inch by 1/4 inch pieces.\u003c/li\u003e\u003cli\u003eBrown bacon on medium high heat in a dutch oven big enough to hold the chicken, about 10 minutes. Remove the cooked bacon, set aside. Keep the bacon fat in the pan. Working in batches if necessary, add onions and chicken, skin side down. Brown the chicken well, on all sides, about 10 minutes. Halfway through the browning, add the garlic and sprinkle the chicken with salt and pepper.\u003c/li\u003e\u003cli\u003eSpoon off any excess fat. Add the chicken stock, wine, and herbs. Add back the bacon. Lower heat to a simmer. Cover and cook for 20 minutes, or until chicken is tender and cooked through. Remove chicken and onions to a separate platter. Remove the bay leaves, herb sprigs, garlic, and discard.\u003c/li\u003e\u003cli\u003eAdd mushrooms to the remaining liquid and turn the heat to high. Boil quickly and reduce the liquid by three fourths until it becomes thick and saucy. Lower the heat, stir in the butter. Return the chicken and onions to the pan to reheat and coat with sauce. Adjust seasoning. Garnish with parsley and serve with potatoes or egg noodles.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"hot",
    "cookingMethod":"multi_cooker",
    "rating":64
  },
  {
    "title":"Blue Cheese Burgers",
    "picture":"https://spoonacular.com/recipeImages/635350-556x370.jpg",
    "ingredients":[
      "kosher salt",
      "ground sirloin",
      "ground sirloin",
      "dry bread crumbs",
      "steak sauce",
      "egg",
      "ground pepper",
      "hamburger buns",
      "blue cheese"
    ],
    "ingredientInfo":{
      "blue cheese":{
        "amount":4,
        "measuresUnit":"ounces",
        "image":"blue-cheese.png",
        "name":"blue cheese"
      },
      "dry bread crumbs":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"breadcrumbs.jpg",
        "name":"dry bread crumbs"
      },
      "egg":{
        "amount":1,
        "measuresUnit":"large",
        "image":"egg.png",
        "name":"egg"
      },
      "ground pepper":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"pepper.jpg",
        "name":"ground pepper"
      },
      "ground sirloin":{
        "amount":0.5,
        "measuresUnit":"pound",
        "image":"fresh-ground-beef.jpg",
        "name":"ground sirloin"
      },
      "hamburger buns":{
        "amount":4,
        "measuresUnit":"",
        "image":"hamburger-bun.jpg",
        "name":"hamburger buns"
      },
      "kosher salt":{
        "amount":0.75,
        "measuresUnit":"tsp",
        "image":"salt.jpg",
        "name":"kosher salt"
      },
      "steak sauce":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"dark-sauce.jpg",
        "name":"steak sauce"
      }
    },
    "summary":"Step 1: In a bowl, add chuck, sirloin, bread crumbs, steak sauce, eggs, salt, and pepper, and mix gently with fork.\nStep 2: Form 4 patties out of this mixture.\nStep 3: Cook the hamburgers patties on each side, then cover with aluminium foil.\nStep 4: Cut the sides of the bun, then grill in a stove-top grill until toasted.\nStep 5: Place patties and a slice of blue cheese in each bun.\nStep 6: Serve hot.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":65
  },
  {
    "title":"White chocolate-cranberry ice cream",
    "picture":"https://spoonacular.com/recipeImages/665203-556x370.jpg",
    "ingredients":[
      "double cream",
      "dried cranberries",
      "egg yolks",
      "full fat milk",
      "sugar",
      "vanilla extract",
      "white chocolate"
    ],
    "ingredientInfo":{
      "double cream":{
        "amount":200,
        "measuresUnit":"ml",
        "image":"white-cream.png",
        "name":"double cream"
      },
      "dried cranberries":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"dried-cranberries.jpg",
        "name":"dried cranberries"
      },
      "egg yolks":{
        "amount":4,
        "measuresUnit":"",
        "image":"egg-yolk.jpg",
        "name":"egg yolks"
      },
      "full fat milk":{
        "amount":300,
        "measuresUnit":"ml",
        "image":"milk.png",
        "name":"full fat milk"
      },
      "sugar":{
        "amount":75,
        "measuresUnit":"g",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      },
      "vanilla extract":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"vanilla-extract.jpg",
        "name":"vanilla extract"
      },
      "white chocolate":{
        "amount":200,
        "measuresUnit":"g",
        "image":"white-chocolate.jpg",
        "name":"white chocolate"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eIn a medium saucepan boil the milk until it starts to bubble gently. Remove from flame and set aside.\u003c/li\u003e\u003cli\u003eMeanwhile whisk the egg yolks and sugar together for 3 minutes until pale yellow and double in volume.\u003c/li\u003e\u003cli\u003ePour the milk on to the egg mixture and stir and then gently pour the mixture back in to the saucepan and simmer very lightly. Stir continuously until the custard mix thickens or coats the back of a wooden spoon.\u003c/li\u003e\u003cli\u003eAdd the chopped chocolate stirring until the heat melts the chocolate.\u003c/li\u003e\u003cli\u003eCool the mixture, stir in the vanilla extract and put it in the fridge to chill.\u003c/li\u003e\u003cli\u003eOnce chilled, add the cream to the mixture. Stir and combine.\u003c/li\u003e\u003cli\u003ePour the mixture into an ice cream machine and churn until frozen.\u003c/li\u003e\u003cli\u003eAdd  dried cranberries in the last 5 minutes of churning and let the machine stir them in.\u003c/li\u003e\u003cli\u003eTransfer to plastic container and place in the freezer for 3 hours before serving.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":26
  },
  {
    "title":"Easy Chicken Tandoori",
    "picture":"https://spoonacular.com/recipeImages/641904-556x370.jpg",
    "ingredients":[
      "chicken",
      "curry paste",
      "fresh coriander leaves",
      "greek yogurt",
      "onions",
      "tomatoes"
    ],
    "ingredientInfo":{
      "chicken":{
        "amount":2,
        "measuresUnit":"pounds",
        "image":"whole-chicken.jpg",
        "name":"chicken"
      },
      "curry paste":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"chili-paste.png",
        "name":"curry paste"
      },
      "fresh coriander leaves":{
        "amount":1,
        "measuresUnit":"leaves",
        "image":"cilantro.png",
        "name":"fresh coriander leaves"
      },
      "greek yogurt":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"plain-yogurt.jpg",
        "name":"greek yogurt"
      },
      "onions":{
        "amount":12,
        "measuresUnit":"servings",
        "image":"brown-onion.png",
        "name":"onions"
      },
      "tomatoes":{
        "amount":12,
        "measuresUnit":"servings",
        "image":"tomato.png",
        "name":"tomatoes"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eMix the tandoori paste and yogurt well. Marinate the chicken pieces well for about 30 minutes\u003c/li\u003e\u003cli\u003ePlace the chicken pieces along with the marinate into a baking pan and bake until golden brown and then turn over  until chicken is done. Keep turning the chicken pieces a few times while baking. \u003c/li\u003e\u003cli\u003eServe with in a bed of rice. Then lay the chicken fillets on the rice and put a bit of greek yogurt\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":57
  },
  {
    "title":"Curried Butternut Squash and Apple Soup",
    "picture":"https://spoonacular.com/recipeImages/641057-556x370.jpg",
    "ingredients":[
      "cauliflower",
      "butternut squash",
      "apple",
      "water",
      "braggs liquid aminos",
      "curry powder",
      "ginger powder",
      "sriracha",
      "tofu"
    ],
    "ingredientInfo":{
      "apple":{
        "amount":2,
        "measuresUnit":"oz",
        "image":"apple.jpg",
        "name":"apple"
      },
      "braggs liquid aminos":{
        "amount":2,
        "measuresUnit":"tsp",
        "image":"dark-sauce.jpg",
        "name":"braggs liquid aminos"
      },
      "butternut squash":{
        "amount":2.5,
        "measuresUnit":"oz",
        "image":"butternut-squash.jpg",
        "name":"butternut squash"
      },
      "cauliflower":{
        "amount":3.5,
        "measuresUnit":"oz",
        "image":"cauliflower.jpg",
        "name":"cauliflower"
      },
      "curry powder":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"curry-powder.jpg",
        "name":"curry powder"
      },
      "ginger powder":{
        "amount":0.25,
        "measuresUnit":"tsp",
        "image":"ginger.png",
        "name":"ginger powder"
      },
      "sriracha":{
        "amount":0.125,
        "measuresUnit":"tsp",
        "image":"hot-sauce-or-tabasco.png",
        "name":"sriracha"
      },
      "tofu":{
        "amount":0.125,
        "measuresUnit":"",
        "image":"tofu.png",
        "name":"tofu"
      },
      "water":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"water.png",
        "name":"water"
      }
    },
    "summary":"Get ready two small saucepans.\nIn one have together 1/2 of the squash and 1/2 of the apple  bring to a simmer, and simmer until just tender when pierced with a fork.\nIn the other, add in the rest of the vegetables along with 1 cup of water  bring this to a rolling boil, reduce to a simmer, and simmer until very very tender. When very tender, remove the vegetables from the heat, and blend these vegetables together (without draining) with the Braggs, spices, sriracha and the additional 1/2 cup water.\nMix the blended part of the soup with the simmered squash and apples. Mix in the cubed tofu, gently re-heat and serve.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":89
  },
  {
    "title":"Kenyan Pilau",
    "picture":"https://spoonacular.com/recipeImages/716344-556x370.jpg",
    "ingredients":[
      "black peppercorns",
      "cardamom pods",
      "cinnamon sticks",
      "cumin seed",
      "garlic",
      "ginger",
      "meat",
      "onion",
      "rice",
      "salt",
      "shrimp",
      "tomatoes",
      "vegetable cooking oil"
    ],
    "ingredientInfo":{
      "black peppercorns":{
        "amount":4,
        "measuresUnit":"",
        "image":"black-pepper.png",
        "name":"black peppercorns"
      },
      "cardamom pods":{
        "amount":4,
        "measuresUnit":"",
        "image":"cardamom.jpg",
        "name":"cardamom pods"
      },
      "cinnamon sticks":{
        "amount":4,
        "measuresUnit":"",
        "image":"cinnamon.jpg",
        "name":"cinnamon sticks"
      },
      "cumin seed":{
        "amount":4,
        "measuresUnit":"teaspoons",
        "image":"ground-cumin.jpg",
        "name":"cumin seed"
      },
      "garlic":{
        "amount":2,
        "measuresUnit":"cloves",
        "image":"garlic.png",
        "name":"garlic"
      },
      "ginger":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"ginger.png",
        "name":"ginger"
      },
      "meat":{
        "amount":1,
        "measuresUnit":"pound",
        "image":"whole-chicken.jpg",
        "name":"meat"
      },
      "onion":{
        "amount":1,
        "measuresUnit":"small",
        "image":"brown-onion.png",
        "name":"onion"
      },
      "rice":{
        "amount":1.5,
        "measuresUnit":"cups",
        "image":"uncooked-white-rice.png",
        "name":"rice"
      },
      "salt":{
        "amount":4,
        "measuresUnit":"servings",
        "image":"salt.jpg",
        "name":"salt"
      },
      "shrimp":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"shrimp.png",
        "name":"shrimp"
      },
      "tomatoes":{
        "amount":2,
        "measuresUnit":"small",
        "image":"tomato.png",
        "name":"tomatoes"
      },
      "vegetable cooking oil":{
        "amount":1,
        "measuresUnit":"",
        "image":"",
        "name":"vegetable cooking oil"
      }
    },
    "summary":"Season and boil the meat and set aside.Chop and crush all the spices and set aside.Heat up the oil and sautГ© the onions till its golden brown. Add the meat and allow to brown a little then add the spices, and the rice and the chopped tomatoes.Add the shrimps and add 2В cups of water and reduce the heat and allow to steam.Consistently check it to make sure the water has dried and the rice is soft. Once the rice is soft, increase the heat and stir it to let all the spices be absorbed in the rice.Serve with Kachumbari or any other salad.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":85
  },
  {
    "title":"Cheesy Chicken Enchilada Quinoa Casserole",
    "picture":"https://spoonacular.com/recipeImages/715421-556x370.jpg",
    "ingredients":[
      "avocado",
      "black pepper",
      "canned black beans",
      "canned tomatoes",
      "chili powder",
      "cooked quinoa",
      "cumin",
      "enchilada sauce",
      "fresh cilantro",
      "green onion tops",
      "roma tomato",
      "salt",
      "shredded cheese",
      "skinless boneless chicken breast",
      "sweet corn",
      "white pepper"
    ],
    "ingredientInfo":{
      "avocado":{
        "amount":1,
        "measuresUnit":"small",
        "image":"avocado.jpg",
        "name":"avocado"
      },
      "black pepper":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"pepper.jpg",
        "name":"black pepper"
      },
      "canned black beans":{
        "amount":15,
        "measuresUnit":"oz",
        "image":"black-beans.jpg",
        "name":"canned black beans"
      },
      "canned tomatoes":{
        "amount":10,
        "measuresUnit":"oz",
        "image":"tomatoes-canned.png",
        "name":"canned tomatoes"
      },
      "chili powder":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"chili-powder.jpg",
        "name":"chili powder"
      },
      "cooked quinoa":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"cooked-quinoa.png",
        "name":"cooked quinoa"
      },
      "cumin":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"ground-cumin.jpg",
        "name":"cumin"
      },
      "enchilada sauce":{
        "amount":10,
        "measuresUnit":"oz",
        "image":"harissa.jpg",
        "name":"enchilada sauce"
      },
      "fresh cilantro":{
        "amount":2,
        "measuresUnit":"Tbsp",
        "image":"cilantro.png",
        "name":"fresh cilantro"
      },
      "green onion tops":{
        "amount":4,
        "measuresUnit":"servings",
        "image":"spring-onions.jpg",
        "name":"green onion tops"
      },
      "roma tomato":{
        "amount":1,
        "measuresUnit":"",
        "image":"roma-tomatoes.png",
        "name":"roma tomato"
      },
      "salt":{
        "amount":4,
        "measuresUnit":"servings",
        "image":"salt.jpg",
        "name":"salt"
      },
      "shredded cheese":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"cheddar-cheese.png",
        "name":"shredded cheese"
      },
      "skinless boneless chicken breast":{
        "amount":1,
        "measuresUnit":"",
        "image":"chicken-breasts.png",
        "name":"skinless boneless chicken breast"
      },
      "sweet corn":{
        "amount":15,
        "measuresUnit":"oz",
        "image":"corn.png",
        "name":"sweet corn"
      },
      "white pepper":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"white-pepper.png",
        "name":"white pepper"
      }
    },
    "summary":"To get started, heat your oven to 350 and prepare an 8x8 baking dish. Cook the quinoa according to the instructions. If you haven't already, boil and shred your chicken breast.In a medium sized mixing bowl add 2 Tbsp cilantro, 1 cup shredded cheese, quinoa, tomatoes with chilis, black beans, half of the can of sweet corn, verde enchilada sauce, cumin, chili powder, white pepper, black pepper, and salt to taste. Mix everything well, and then pour into the 8 x8 baking dish. Cover with the last cup of cheese and then bake in the oven for 15 minutes.Remove from oven and allow to cool for 5 minutes. Then top with chopped Roma tomato, chopped avocado, green onion tops, and remaining 2 tsp of cilantro. Serve immediately.",
    "readyInMinutes":30,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":97
  },
  {
    "title":"Sausage \u0026 Pepperoni Stromboli",
    "picture":"https://spoonacular.com/recipeImages/776505-556x370.jpg",
    "ingredients":[
      "egg",
      "italian sausage",
      "parmesan cheese",
      "pepperoni",
      "pizza dough",
      "pizza sauce",
      "shredded mozzarella cheese"
    ],
    "ingredientInfo":{
      "egg":{
        "amount":1,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"egg"
      },
      "italian sausage":{
        "amount":1,
        "measuresUnit":"lb",
        "image":"raw-pork-sausage.png",
        "name":"italian sausage"
      },
      "parmesan cheese":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"parmesan.jpg",
        "name":"parmesan cheese"
      },
      "pepperoni":{
        "amount":1,
        "measuresUnit":"package",
        "image":"pepperoni.png",
        "name":"pepperoni"
      },
      "pizza dough":{
        "amount":1,
        "measuresUnit":"ball",
        "image":"pizza-dough.jpg",
        "name":"pizza dough"
      },
      "pizza sauce":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"tomato-sauce-or-pasta-sauce.jpg",
        "name":"pizza sauce"
      },
      "shredded mozzarella cheese":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"shredded-cheese-white.jpg",
        "name":"shredded mozzarella cheese"
      }
    },
    "summary":"\u003cp\u003eCook Italian sausage in a medium skillet until browned. Drain on paper towels and crumble into small pieces.\u003c/p\u003e\u003cp\u003eHeat oven to 450 degrees.\u003c/p\u003e\u003cp\u003ePlace parchment paper on a baking pan and lightly flour. Roll dough out to form a large rectangle.\u003c/p\u003e\u003cp\u003eBrush a thin layer of pizza sauce on dough. Layer on crumbled sausage then mozzarella cheese.\u003c/p\u003e\u003cp\u003ePlace pepperoni slices on top of cheese then sprinkle with 3/4 cup Parmesan cheese.\u003c/p\u003e\u003cp\u003eGently roll in short sides of dough to form a seal around filling.\u003c/p\u003e\u003cp\u003eStarting on long side of dough, carefully roll into one third of Stromboli. Using parchment, roll dough again until reaching opposite end. Pinch together and place seam side down.\u003c/p\u003e\u003cp\u003eMix egg with 1 tablespoon of water. Brush egg mixture over Stromboli and sprinkle with remaining 1/4 Parmesan cheese.\u003c/p\u003e\u003cp\u003ePlace in oven and IMMEDIATELY TURN OVEN DOWN to 350 degrees.\u003c/p\u003e\u003cp\u003eBake for approximately 18 minutes. Remove from oven and let rest for 10 minutes. \u003c/p\u003e\u003cp\u003eTest center of turnover. If dough is not completely cooked, slice Stromboli in half and bake for additional 5-8 minutes.\u003c/p\u003e\u003cp\u003eServe with warm pizza sauce\u003c/p\u003e",
    "readyInMinutes":28,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":68
  },
  {
    "title":"Maple-Glazed Apple Cookies",
    "picture":"https://spoonacular.com/recipeImages/650939-556x370.jpg",
    "ingredients":[
      "apple",
      "baking soda",
      "brown sugar",
      "butter",
      "cinnamon",
      "egg",
      "flour",
      "ground cloves",
      "maple flavoring",
      "milk",
      "nutmeg",
      "powdered sugar",
      "salt"
    ],
    "ingredientInfo":{
      "apple":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"apple.jpg",
        "name":"apple"
      },
      "baking soda":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"white-powder.jpg",
        "name":"baking soda"
      },
      "brown sugar":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"light-brown-sugar.jpg",
        "name":"brown sugar"
      },
      "butter":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"butter-sliced.jpg",
        "name":"butter"
      },
      "cinnamon":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"cinnamon.jpg",
        "name":"cinnamon"
      },
      "egg":{
        "amount":1,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"egg"
      },
      "flour":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"flour.png",
        "name":"flour"
      },
      "ground cloves":{
        "amount":0.25,
        "measuresUnit":"teaspoon",
        "image":"cloves.jpg",
        "name":"ground cloves"
      },
      "maple flavoring":{
        "amount":0.25,
        "measuresUnit":"teaspoon",
        "image":"vanilla-extract.jpg",
        "name":"maple flavoring"
      },
      "milk":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"milk.png",
        "name":"milk"
      },
      "nutmeg":{
        "amount":0.25,
        "measuresUnit":"teaspoon",
        "image":"ground-nutmeg.jpg",
        "name":"nutmeg"
      },
      "powdered sugar":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"powdered-sugar.jpg",
        "name":"powdered sugar"
      },
      "salt":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"salt.jpg",
        "name":"salt"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePreheat oven to 400 degrees F.\u003c/li\u003e\u003cli\u003eCream 1/2 cup butter and sugar together in a large mixing bowl until light in color, about 1-2 minutes. Add egg and 1/4 cup milk; stir to combine. Add remaining dry ingredients (except apple) and stir just until combined. Gently fold in the apples.  Drop by heaping tablespoons onto a greased baking sheet.\u003c/li\u003e\u003cli\u003eBake at 400 for about 10-12 min or until tops don't look wet anymore.\u003c/li\u003e\u003cli\u003eTo make the glaze, whisk together 1 tablespoon melted butter, 1 cup powdered sugar, 1/4 tsp maple flavoring, and 3 tablespoons milk in a small bowl. When the cookies have cooled for a few minutes, drizzle each cookie with glaze.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":15
  },
  {
    "title":"Coconut Macaroons With Chocolate Drizzle",
    "picture":"https://spoonacular.com/recipeImages/639779-556x370.jpg",
    "ingredients":[
      "flour",
      "sweetened coconut",
      "salt",
      "vanilla extract",
      "sweetened condensed milk"
    ],
    "ingredientInfo":{
      "flour":{
        "amount":0.6666666666666666,
        "measuresUnit":"cup",
        "image":"flour.png",
        "name":"flour"
      },
      "salt":{
        "amount":0.25,
        "measuresUnit":"teaspoon",
        "image":"salt.jpg",
        "name":"salt"
      },
      "sweetened coconut":{
        "amount":5.5,
        "measuresUnit":"cups",
        "image":"shredded-coconut.jpg",
        "name":"sweetened coconut"
      },
      "sweetened condensed milk":{
        "amount":14,
        "measuresUnit":"oz",
        "image":"evaporated-milk.png",
        "name":"sweetened condensed milk"
      },
      "vanilla extract":{
        "amount":2,
        "measuresUnit":"teaspoons",
        "image":"vanilla-extract.jpg",
        "name":"vanilla extract"
      }
    },
    "summary":"Combine the flour, coconut and salt in a large bowl.\nIn a smaller bowl, combine the vanilla and the can of sweetened condensed milk and mix well.\nAdd this goo to the dry ingredients and mix with a wooden spoon, or you could use your hands (sounds rather messy to me). I probably wouldn't use a mixer unless it's on a super low speed. This batter is going to be THICK.\nLine baking sheets with parchment paper, and, using a big spoon or ice cream scooper, scoop the batter/dough onto the sheets.\nIn a preheated 350 degree oven, bake the macaroons for about 20 minutes, or until golden/toasty looking.\nDrizzle some melted semi-sweet chocolate on top or use a chocolate/baker's chocolate mix and go for the dipped variety.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":20
  },
  {
    "title":"Moroccan Lemon Shish Kebabs",
    "picture":"https://spoonacular.com/recipeImages/652433-556x370.jpg",
    "ingredients":[
      "boneless skinless chicken breast fillets",
      "lemon",
      "parsley",
      "fresh rosemary leaves",
      "fresh thyme leaves",
      "garlic",
      "black peppercorns",
      "juice of lemon",
      "olive oil"
    ],
    "ingredientInfo":{
      "black peppercorns":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"black-pepper.png",
        "name":"black peppercorns"
      },
      "boneless skinless chicken breast fillets":{
        "amount":1,
        "measuresUnit":"pound",
        "image":"chicken-breasts.png",
        "name":"boneless skinless chicken breast fillets"
      },
      "fresh rosemary leaves":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"rosemary.jpg",
        "name":"fresh rosemary leaves"
      },
      "fresh thyme leaves":{
        "amount":2,
        "measuresUnit":"teaspoons",
        "image":"thyme.jpg",
        "name":"fresh thyme leaves"
      },
      "garlic":{
        "amount":2,
        "measuresUnit":"cloves",
        "image":"garlic.png",
        "name":"garlic"
      },
      "juice of lemon":{
        "amount":1,
        "measuresUnit":"",
        "image":"lemon-juice.jpg",
        "name":"juice of lemon"
      },
      "lemon":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"lemon.png",
        "name":"lemon"
      },
      "olive oil":{
        "amount":2,
        "measuresUnit":"teaspoons",
        "image":"olive-oil.jpg",
        "name":"olive oil"
      },
      "parsley":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"parsley.jpg",
        "name":"parsley"
      }
    },
    "summary":"The chicken breast fillets trimmed of fat, cut into 1\" cubes\nFor Moroccan lemon marinade\nPut everything in a bowl\nAdd this Marinated chicken cubes, mix well, cover with Plastic Wrap and let stand at least 1/2 hour before use ... I like to leave more time in the fridge\nThe time for rest, place the bamboo skewers or metal rods in a tray with water (to cover only), that is to avoid burning\nThread the meat on each rod are more or less as 4 pieces of meat\nShish Kebbab cooking on the grill, skillet or the oven ... I like the first option, you know the smell of charcoal, wood ...\nOnce the Shish Kebab cooked, serve immediately on a bed of Rice, Mediterranean Cous Cous, etc. ... or only with Pita Bread and ready",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":37
  },
  {
    "title":"Stir Fried Cabbage and Tomatoes",
    "picture":"https://spoonacular.com/recipeImages/661653-556x370.jpg",
    "ingredients":[
      "cabbage",
      "tomatoes",
      "garlic",
      "salt",
      "sugar",
      "ketchup",
      "water"
    ],
    "ingredientInfo":{
      "cabbage":{
        "amount":0.5,
        "measuresUnit":"",
        "image":"cabbage.jpg",
        "name":"cabbage"
      },
      "garlic":{
        "amount":2,
        "measuresUnit":"",
        "image":"garlic.png",
        "name":"garlic"
      },
      "ketchup":{
        "amount":1,
        "measuresUnit":"tbsp",
        "image":"ketchup.png",
        "name":"ketchup"
      },
      "salt":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"salt.jpg",
        "name":"salt"
      },
      "sugar":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      },
      "tomatoes":{
        "amount":2,
        "measuresUnit":"medium",
        "image":"tomato.png",
        "name":"tomatoes"
      },
      "water":{
        "amount":2,
        "measuresUnit":"tbsp",
        "image":"water.png",
        "name":"water"
      }
    },
    "summary":"Coarsely shred the cabbage and rinse well. Drain and set aside.\nRinse tomatoes and dice, set aside.\nHeat wok with some oil to saute the garlic. Put in the cabbage and cook until soft, add a little water if necessary. Add diced tomatoes and seasoning with 2-3 tablespoons of water into it. Stir fry to combine and serve immediately.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":90
  },
  {
    "title":"No-Bake Chocolate Peanut Butter Pie",
    "picture":"https://spoonacular.com/recipeImages/653185-556x370.jpg",
    "ingredients":[
      "cream cheese",
      "creamy peanut butter",
      "dark chocolate chips",
      "peanut butter cups",
      "powdered sugar",
      "roasted peanuts",
      "soy buttery spread",
      "vanilla yogurt",
      "whipped topping"
    ],
    "ingredientInfo":{
      "cream cheese":{
        "amount":2,
        "measuresUnit":"tablespoons",
        "image":"cream-cheese.jpg",
        "name":"cream cheese"
      },
      "creamy peanut butter":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"peanut-butter.png",
        "name":"creamy peanut butter"
      },
      "dark chocolate chips":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"chocolate-chips.jpg",
        "name":"dark chocolate chips"
      },
      "peanut butter cups":{
        "amount":10,
        "measuresUnit":"",
        "image":"peanut-butter-cup.jpg",
        "name":"peanut butter cups"
      },
      "powdered sugar":{
        "amount":1.25,
        "measuresUnit":"cups",
        "image":"powdered-sugar.jpg",
        "name":"powdered sugar"
      },
      "roasted peanuts":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"peanuts.png",
        "name":"roasted peanuts"
      },
      "soy buttery spread":{
        "amount":5,
        "measuresUnit":"tablespoons",
        "image":"light-buttery-spread.png",
        "name":"soy buttery spread"
      },
      "vanilla yogurt":{
        "amount":3,
        "measuresUnit":"tablespoons",
        "image":"vanilla-yogurt.png",
        "name":"vanilla yogurt"
      },
      "whipped topping":{
        "amount":1.5,
        "measuresUnit":"containers",
        "image":"whipped-cream.jpg",
        "name":"whipped topping"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eCRUST:\u003c/li\u003e\u003cli\u003eCrush the cookies in a food processor until fine crumbs.\u003c/li\u003e\u003cli\u003eTransfer to a small bowl and add the earth balance.\u003c/li\u003e\u003cli\u003eCombine with a fork, or hands until fully blended and begins to take shape.\u003c/li\u003e\u003cli\u003ePress evenly into a 9 inch pie plate. My crust ended at about 1/2 inch before the rim.\u003c/li\u003e\u003cli\u003eChill in the freezer while preparing the filling.\u003c/li\u003e\u003cli\u003eFILLING:\u003c/li\u003e\u003cli\u003eIn a stand mixer set with a paddle attachment, or using a hand held mixer set on medium speed, beat peanut butter with the cream cheese.\u003c/li\u003e\u003cli\u003eAfter about a minute add the coconut yogurt.\u003c/li\u003e\u003cli\u003eOnce combined, reduce to low speed and add sifted powdered sugar, increase back to medium speed and beat until smooth.\u003c/li\u003e\u003cli\u003eTurn off mixer, then add in 1/2 cup peanuts, and broken pieces of peanut butter cups and combine well.\u003c/li\u003e\u003cli\u003eFold in thawed whipped topping until well blended.\u003c/li\u003e\u003cli\u003ePour the filling into the chilled crust and spread evenly.\u003c/li\u003e\u003cli\u003eSprinkle the 3 tablespoons of peanuts, dark chocolate chips, and the reserved cookie crumbs on top.\u003c/li\u003e\u003cli\u003eChill for at least another 90 minutes before serving.\u003c/li\u003e\u003cli\u003eSlice with a sharp knife to pierce through all the peanut butter chunks in the pie.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":43
  },
  {
    "title":"Bacon \u0026 Potato Soup, Gluten \u0026 Dairy Free",
    "picture":"https://spoonacular.com/recipeImages/633263-556x370.jpg",
    "ingredients":[
      "olive oil",
      "lean bacon",
      "yellow onions",
      "garlic cloves",
      "chicken stock",
      "baking mix",
      "savoy cabbage",
      "worcestershire sauce",
      "dijon mustard",
      "flat leaf parsley",
      "garnish"
    ],
    "ingredientInfo":{
      "baking mix":{
        "amount":2.25,
        "measuresUnit":"cups",
        "image":"buttermilk-biscuits.jpg",
        "name":"baking mix"
      },
      "chicken stock":{
        "amount":7,
        "measuresUnit":"cups",
        "image":"chicken-broth.png",
        "name":"chicken stock"
      },
      "dijon mustard":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"dijon-mustard.jpg",
        "name":"dijon mustard"
      },
      "flat leaf parsley":{
        "amount":3,
        "measuresUnit":"Tbsp",
        "image":"parsley.jpg",
        "name":"flat leaf parsley"
      },
      "garlic cloves":{
        "amount":4,
        "measuresUnit":"",
        "image":"garlic.png",
        "name":"garlic cloves"
      },
      "garnish":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"",
        "name":"garnish"
      },
      "lean bacon":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"raw-bacon.png",
        "name":"lean bacon"
      },
      "olive oil":{
        "amount":2,
        "measuresUnit":"Tbsp",
        "image":"olive-oil.jpg",
        "name":"olive oil"
      },
      "savoy cabbage":{
        "amount":3,
        "measuresUnit":"cups",
        "image":"savoy-cabbage.jpg",
        "name":"savoy cabbage"
      },
      "worcestershire sauce":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"dark-sauce.jpg",
        "name":"worcestershire sauce"
      },
      "yellow onions":{
        "amount":2,
        "measuresUnit":"",
        "image":"brown-onion.png",
        "name":"yellow onions"
      }
    },
    "summary":"In a large stock pot, heat the olive oil over medium-high heat. Add the chopped bacon and onions, cooking until the bacon is crisp and the onions are translucent. Add garlic and cook for one more minute, until fragrant.\nAdd the chicken or vegetable stock, potatoes, cabbage, Worcestershire sauce and mustard. Mix well and season with salt and pepper. Bring soup to a boil, then reduce heat and simmer for 30 minutes, or until potatoes are tender but not disintegrating.\nRemove the pot from heat and allow it to cool for 5 minutes. Transfer 2 1/2 cups of the soup to a blender and pulse quickly to achiever a coarse puree. Pour blended portion back into the stock pot and return to heat. Cook, stirring often, for 5-10 minutes or until heated through.\nStir in parsley and ladle into serving bowls. Serve with a wedge of lemon and gluten free garlic croutons.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":70
  },
  {
    "title":"BLT Sandwich",
    "picture":"https://spoonacular.com/recipeImages/635342-556x370.jpg",
    "ingredients":[
      "bell pepper",
      "bread",
      "lettuce",
      "mayonnaise",
      "thick-cut bacon",
      "tomato"
    ],
    "ingredientInfo":{
      "bell pepper":{
        "amount":2,
        "measuresUnit":"servings",
        "image":"yellow-bell-pepper.jpg",
        "name":"bell pepper"
      },
      "bread":{
        "amount":4,
        "measuresUnit":"pieces",
        "image":"white-bread.jpg",
        "name":"bread"
      },
      "lettuce":{
        "amount":2,
        "measuresUnit":"servings",
        "image":"iceberg-lettuce.jpg",
        "name":"lettuce"
      },
      "mayonnaise":{
        "amount":2,
        "measuresUnit":"servings",
        "image":"mayonnaise.png",
        "name":"mayonnaise"
      },
      "thick-cut bacon":{
        "amount":8,
        "measuresUnit":"pieces",
        "image":"raw-bacon.png",
        "name":"thick-cut bacon"
      },
      "tomato":{
        "amount":1,
        "measuresUnit":"",
        "image":"tomato.png",
        "name":"tomato"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eBrown bacon in a skillet\u003c/li\u003e\u003cli\u003eRemove and pat off excess oil\u003c/li\u003e\u003cli\u003eSlice tomato into 1/4 inches slices\u003c/li\u003e\u003cli\u003eToast bread\u003c/li\u003e\u003cli\u003eSpread a thin layer of mayonnaise on bread\u003c/li\u003e\u003cli\u003eLayer all ingredients on bread and close sandwich\u003c/li\u003e\u003cli\u003eAdd fresh cracked black pepper\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":72
  },
  {
    "title":"Brown sugar \u0026 Spice Sugar cookie frogs",
    "picture":"https://spoonacular.com/recipeImages/636300-556x370.jpg",
    "ingredients":[
      "unbleached flour",
      "baking powder",
      "cinnamon",
      "ginger",
      "nutmeg",
      "allspice",
      "granulated sugar",
      "light brown sugar",
      "butter",
      "egg",
      "vanilla extract"
    ],
    "ingredientInfo":{
      "allspice":{
        "amount":0.125,
        "measuresUnit":"tsp",
        "image":"allspice-ground.jpg",
        "name":"allspice"
      },
      "baking powder":{
        "amount":2,
        "measuresUnit":"tsp",
        "image":"white-powder.jpg",
        "name":"baking powder"
      },
      "butter":{
        "amount":2,
        "measuresUnit":"sticks",
        "image":"butter-sliced.jpg",
        "name":"butter"
      },
      "cinnamon":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"cinnamon.jpg",
        "name":"cinnamon"
      },
      "egg":{
        "amount":1,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"egg"
      },
      "ginger":{
        "amount":0.25,
        "measuresUnit":"tsp",
        "image":"ginger.png",
        "name":"ginger"
      },
      "granulated sugar":{
        "amount":0.5,
        "measuresUnit":"c",
        "image":"sugar-in-bowl.png",
        "name":"granulated sugar"
      },
      "light brown sugar":{
        "amount":0.5,
        "measuresUnit":"c",
        "image":"light-brown-sugar.jpg",
        "name":"light brown sugar"
      },
      "nutmeg":{
        "amount":0.125,
        "measuresUnit":"tsp",
        "image":"ground-nutmeg.jpg",
        "name":"nutmeg"
      },
      "unbleached flour":{
        "amount":3,
        "measuresUnit":"c",
        "image":"flour.png",
        "name":"unbleached flour"
      },
      "vanilla extract":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"vanilla-extract.jpg",
        "name":"vanilla extract"
      }
    },
    "summary":"Whisk the flour, baking powder and spices, set aside.\nCream the sugar and butter. Add the egg and extracts and mix until well-blended.\nGradually add the flour mixture and beat just until combined, scraping down the bowl, especially the bottom. (The dough will be quite thick...you may need to knead in stray bits of flour from the bottom of the bowl by hand.)\nRoll on a floured surface and cut into shapes. Place on parchment lined baking sheets and bake for 9-12 minutes, depending on the size of your cutter. Let sit a few minutes on the sheet, then transfer to a cooling rack.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":9
  },
  {
    "title":"Breakfast: Waffles",
    "picture":"https://spoonacular.com/recipeImages/636087-556x370.jpg",
    "ingredients":[
      "baking powder",
      "butter",
      "eggs",
      "flour",
      "salt",
      "sugar",
      "sugar",
      "vanilla extract",
      "whole milk"
    ],
    "ingredientInfo":{
      "baking powder":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"white-powder.jpg",
        "name":"baking powder"
      },
      "butter":{
        "amount":0.25,
        "measuresUnit":"c",
        "image":"butter-sliced.jpg",
        "name":"butter"
      },
      "eggs":{
        "amount":3,
        "measuresUnit":"large",
        "image":"egg.png",
        "name":"eggs"
      },
      "flour":{
        "amount":1.75,
        "measuresUnit":"cups",
        "image":"flour.png",
        "name":"flour"
      },
      "salt":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"salt.jpg",
        "name":"salt"
      },
      "sugar":{
        "amount":4,
        "measuresUnit":"servings",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      },
      "vanilla extract":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"vanilla-extract.jpg",
        "name":"vanilla extract"
      },
      "whole milk":{
        "amount":1.5,
        "measuresUnit":"cups",
        "image":"milk.png",
        "name":"whole milk"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePreheat the waffle iron. In a large bowl add the flour, baking powder, sugar and salt. Mix well together.\u003c/li\u003e\u003cli\u003eIn another bowl add the well beaten eggs, melted butter, vanilla extract and milk.\u003c/li\u003e\u003cli\u003eAdd the liquid ingredients to the dry ingredients and gently whisk together well.\u003c/li\u003e\u003cli\u003eSpoon  cup in the center of the hot waffle iron, or amount recommended by manufacturer. Spread the batter  away from the edge of the iron. Close the lid and cook until the waffle is nice and golden brown.  Serve with sprinkled powder sugar.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":49
  },
  {
    "title":"Salsa Verde Chicken Tamales",
    "picture":"https://spoonacular.com/recipeImages/659180-556x370.jpg",
    "ingredients":[
      "roasted chicken",
      "chicken",
      "garlic",
      "onion",
      "cumin",
      "sea salt",
      "olive oil",
      "water",
      "salsa verde",
      "tomatillos",
      "chile",
      "chile",
      "chiles",
      "yellow onion",
      "cilantro",
      "sea salt",
      "filo dough",
      "masa flour",
      "masa flour",
      "sea salt",
      "palm oil",
      "chicken broth",
      "salsa verde"
    ],
    "ingredientInfo":{
      "chicken":{
        "amount":4,
        "measuresUnit":"lb",
        "image":"whole-chicken.jpg",
        "name":"chicken"
      },
      "chicken broth":{
        "amount":4,
        "measuresUnit":"cups",
        "image":"chicken-broth.png",
        "name":"chicken broth"
      },
      "chile":{
        "amount":1,
        "measuresUnit":"",
        "image":"red-chili.jpg",
        "name":"chile"
      },
      "chiles":{
        "amount":2,
        "measuresUnit":"",
        "image":"red-chili.jpg",
        "name":"chiles"
      },
      "cilantro":{
        "amount":0.3333333333333333,
        "measuresUnit":"",
        "image":"cilantro.png",
        "name":"cilantro"
      },
      "cumin":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"ground-cumin.jpg",
        "name":"cumin"
      },
      "filo dough":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"filo-dough.png",
        "name":"filo dough"
      },
      "garlic":{
        "amount":10,
        "measuresUnit":"cloves",
        "image":"garlic.png",
        "name":"garlic"
      },
      "masa flour":{
        "amount":6,
        "measuresUnit":"cups",
        "image":"corn-flour.jpg",
        "name":"masa flour"
      },
      "olive oil":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"olive-oil.jpg",
        "name":"olive oil"
      },
      "onion":{
        "amount":0.5,
        "measuresUnit":"large",
        "image":"brown-onion.png",
        "name":"onion"
      },
      "palm oil":{
        "amount":6,
        "measuresUnit":"Tbs",
        "image":"vegetable-oil.jpg",
        "name":"palm oil"
      },
      "roasted chicken":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"rotisserie-chicken.png",
        "name":"roasted chicken"
      },
      "salsa verde":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"salsa-verde.png",
        "name":"salsa verde"
      },
      "sea salt":{
        "amount":2,
        "measuresUnit":"tsp",
        "image":"salt.jpg",
        "name":"sea salt"
      },
      "tomatillos":{
        "amount":10,
        "measuresUnit":"",
        "image":"tomatillos.jpg",
        "name":"tomatillos"
      },
      "water":{
        "amount":4,
        "measuresUnit":"cup",
        "image":"water.png",
        "name":"water"
      },
      "yellow onion":{
        "amount":0.25,
        "measuresUnit":"",
        "image":"brown-onion.png",
        "name":"yellow onion"
      }
    },
    "summary":"Roasted Chicken\n1. Place your chicken into a roasting pan or dutch oven. I left the organs and the neck for added flavor.\n2. Put garlic cloves under the skin of the chicken\n3. Chop the onion and scatter around the bottom of the pan. Pour in water.\nSalsa Verde\n1. Remove the hulls from the tomatillos. They will be sticky, don't worry if you can't wash it all off. Remove the stems.\n2. Place all the tomatillos and chiles into a sauce pot. Cover with water and boil until soft.\n3. Remove the tomatillos and chiles from the water and place in a blender. Add all the remaining ingredients and blend until smooth.\n4. It's as simple as that. A yummy salsa perfect for tamales, eating with chip or putting on top of whatever you're eating!\nTamale Dough (Masa)\nMix all together to make a soft, sticky dough.\nSalsa Verde Tamales\nTamale Dough\nShreded chicken\n2 cups Salsa Verde\n10 to 40 green olives\n4 medium potatoes\n1. Remove the corn husks from the bag and place them in a pot of hot water. Continue to heat on low for   at least 30 minutes. You want them to be pliable.\n2. Add your shreded chicken to a skillet and cover with  about two cups of salsa verde.\n3. If you have large olives you will want to cut them in 1/2 or in 1/4. If they are small you might just want to leave them whole. I prefer to remove the pit but you don't have too, just watch out when you bit into the tamale. Place in small bowl and set aside.\n4. Cut the potatoes into small, thick strips. You don't have to remove the peel.  I used red potatoes because that is all I had. Normally I would use  regular old potatoes. Place in a small bowl, cover with water  to previent them from turning brown and set aside.\n5. Now comes the fun part! Asembling the tamales! It is best to do this with a firend. My sister helped me since I am visiting her right now. Remove the soaked corn husks from the pot.\n6. Pick out the husk you want to use. Normally one big one is enough. If the are small you may have to use two, overlapping a bit.\n7. Spred about 2TBS of the dough across the bottom half of the husk, narrow end pointing away from you about 1/8 in thick. Going a little more than half way up, leaving a tiny bit of room on the sides.\n8. Now place a small amount of chicken in the middle of the dough. I can't remember the exact amount I used.\n9. Place a potatoe sclice in the middle of the chicken and an olive at the end.\n10. Fold in one side and then the other. Fold up the bottom. If you want you can tie them closed with a small strip of husk. I don't.\n11. In a large pot pour in a about two cups of water. Place as many tamales that will fit in the pot, open end facing up. Place the lid on the pot and bring to a boil. Once boiling, turn down the heat and steam for 30 to 45 minutes.\n12. Remove from the pot and enjoy!! You will know that they are done when the husk can be pulled away from the masa and not stick. I served mine with refried beans and rice. The rice recipe will soon to follow, hopefully! This recipe made almost 4 dozen tamales. They can be frozen to save for later. Simply let them defrost and steam until warm.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":25
  },
  {
    "title":"Eggplant Fries with Tzatziki Sauce",
    "picture":"https://spoonacular.com/recipeImages/642287-556x370.jpg",
    "ingredients":[
      "eggplants",
      "bread crumbs",
      "seasoning mix",
      "low fat plain yogurt",
      "egg"
    ],
    "ingredientInfo":{
      "bread crumbs":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"breadcrumbs.jpg",
        "name":"bread crumbs"
      },
      "egg":{
        "amount":1,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"egg"
      },
      "eggplants":{
        "amount":0.5,
        "measuresUnit":"",
        "image":"eggplant.png",
        "name":"eggplants"
      },
      "low fat plain yogurt":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"plain-yogurt.jpg",
        "name":"low fat plain yogurt"
      },
      "seasoning mix":{
        "amount":1,
        "measuresUnit":"serving",
        "image":"seasoning.png",
        "name":"seasoning mix"
      }
    },
    "summary":"-Preheat oven to 450F\n-Mix salt, garlic powder, italian seasonings, and paprika in a bowl.\n- In another bowl, mix yogurt \u0026 egg together.\n-First place the eggplant strips into egg/yogurt mixture then coat them in flour mixture evenly.\n-Place them in a parchment paper or greased baking pan and bake for 10-15 minutes rotating once and until slightly brown.\nFor the tzatziki sauce:\n- In a food processor, put cucumber (seeded, peeled), yogurt, dill, lemon juice, garlic cloves, salt together until smooth.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":66
  },
  {
    "title":"Dulce De Leche CrГЁme BrГ»lГ©e",
    "picture":"https://spoonacular.com/recipeImages/641727-556x370.jpg",
    "ingredients":[
      "cream",
      "dulce de leche",
      "sugar",
      "egg yolks"
    ],
    "ingredientInfo":{
      "cream":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"fluid-cream.jpg",
        "name":"cream"
      },
      "dulce de leche":{
        "amount":3,
        "measuresUnit":"tablespoons",
        "image":"dulce-de-leche.png",
        "name":"dulce de leche"
      },
      "egg yolks":{
        "amount":5,
        "measuresUnit":"large",
        "image":"egg-yolk.jpg",
        "name":"egg yolks"
      },
      "sugar":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      }
    },
    "summary":"Oven: 325F\nPlace six ramekins in a water bath. Whisk eggs and sugar until pale, then slowly pour the hot cream into the yolks, whisking thoroughly. Pour custards into ramekins and bake ~35 min. until set. Chill at least 3 hours before serving.\nTo serve, sprinkle ~2 tsp sugar evenly over each custard and heat with a kitchen torch until a burnt crust forms atop each custard.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":83
  },
  {
    "title":"Brown Sugar Sponge Cookie w/ Chocolate Covered Caramels",
    "picture":"https://spoonacular.com/recipeImages/636315-556x370.jpg",
    "ingredients":[
      "all purpose flour",
      "baking powder",
      "dark brown sugar",
      "eggs",
      "lemon zest",
      "salt",
      "vanilla extract",
      "rolos"
    ],
    "ingredientInfo":{
      "all purpose flour":{
        "amount":0.3333333333333333,
        "measuresUnit":"cup",
        "image":"flour.png",
        "name":"all purpose flour"
      },
      "baking powder":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"white-powder.jpg",
        "name":"baking powder"
      },
      "dark brown sugar":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"dark-brown-sugar.png",
        "name":"dark brown sugar"
      },
      "eggs":{
        "amount":2,
        "measuresUnit":"large",
        "image":"egg.png",
        "name":"eggs"
      },
      "lemon zest":{
        "amount":0.125,
        "measuresUnit":"teaspoon",
        "image":"zest-lemon.jpg",
        "name":"lemon zest"
      },
      "rolos":{
        "amount":24,
        "measuresUnit":"",
        "image":"",
        "name":"rolos"
      },
      "salt":{
        "amount":0.125,
        "measuresUnit":"teaspoon",
        "image":"salt.jpg",
        "name":"salt"
      },
      "vanilla extract":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"vanilla-extract.jpg",
        "name":"vanilla extract"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePreheat the oven to 375 degrees. Line baking sheets with parchment paper and spray with nonstick cooking spray.\u003c/li\u003e\u003cli\u003eIn the bowl of a stand mixer on medium speed, beat the egg whites until foamy. Add the baking powder, salt and sugar and continue beating on medium speed until stiff peaks form.\u003c/li\u003e\u003cli\u003eIn another bowl combine egg yolks, vanilla and lemon zest, beat with fork until thoroughly mixed. Fold the yolks into the egg white mixture just until combined.\u003c/li\u003e\u003cli\u003eSift the flour over the egg mixture and fold in until the batter is smooth and light. Drop just about 2 teaspoons of batter for each cookie about 2 inches apart onto baking sheets. Top each off with a Rolo in the center and press down lightly.\u003c/li\u003e\u003cli\u003eBake for 12 minutes or until golden. Cool on the sheets for 5 minutes then move to racks to cool completely.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":10
  },
  {
    "title":"Guacamole",
    "picture":"https://spoonacular.com/recipeImages/645988-556x370.jpg",
    "ingredients":[
      "haas avocados",
      "onion",
      "jalapeno pepper",
      "fresh cilantro",
      "salt",
      "black pepper",
      "lime juice",
      "fresh basil",
      "oregano"
    ],
    "ingredientInfo":{
      "black pepper":{
        "amount":0.125,
        "measuresUnit":"teaspoon",
        "image":"pepper.jpg",
        "name":"black pepper"
      },
      "fresh basil":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"basil.jpg",
        "name":"fresh basil"
      },
      "fresh cilantro":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"cilantro.png",
        "name":"fresh cilantro"
      },
      "haas avocados":{
        "amount":2,
        "measuresUnit":"large",
        "image":"avocado.jpg",
        "name":"haas avocados"
      },
      "jalapeno pepper":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"jalapeno-pepper.png",
        "name":"jalapeno pepper"
      },
      "lime juice":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"lime-juice.png",
        "name":"lime juice"
      },
      "onion":{
        "amount":2,
        "measuresUnit":"tablespoons",
        "image":"brown-onion.png",
        "name":"onion"
      },
      "oregano":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"oregano.jpg",
        "name":"oregano"
      },
      "salt":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"salt.jpg",
        "name":"salt"
      }
    },
    "summary":"Cut the avocados in half lengthwise. Drive chef's knife into large pit and twist to remove.  Scoop the avocado out from the skin into a bowl.\nMash the avocado with a fork, add in onion and cilantro, oregano and basil..  Mix to incorporate.\nSeason to taste with salt, pepper and lime juice.\nServe immediately or chill.  If you are not going to eat the guacamole immediately, store in refrigerator with plastic film wrap pressed onto the guacamole.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":80
  },
  {
    "title":"Lemon Coconut Granola",
    "picture":"https://spoonacular.com/recipeImages/649567-556x370.jpg",
    "ingredients":[
      "coconut oil",
      "desiccated coconut",
      "dried currants",
      "flax seeds",
      "honey",
      "juice of lemon",
      "lemon peel",
      "nuts",
      "rolled oats"
    ],
    "ingredientInfo":{
      "coconut oil":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"oil-coconut.jpg",
        "name":"coconut oil"
      },
      "desiccated coconut":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"shredded-coconut.jpg",
        "name":"desiccated coconut"
      },
      "dried currants":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"blueberries-dried.jpg",
        "name":"dried currants"
      },
      "flax seeds":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"flax-seeds.png",
        "name":"flax seeds"
      },
      "honey":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"honey.png",
        "name":"honey"
      },
      "juice of lemon":{
        "amount":3,
        "measuresUnit":"",
        "image":"lemon-juice.jpg",
        "name":"juice of lemon"
      },
      "lemon peel":{
        "amount":1,
        "measuresUnit":"",
        "image":"zest-lemon.jpg",
        "name":"lemon peel"
      },
      "nuts":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"nuts-mixed.jpg",
        "name":"nuts"
      },
      "rolled oats":{
        "amount":3,
        "measuresUnit":"cups",
        "image":"rolled-oats.jpg",
        "name":"rolled oats"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePreheat oven to 150C.\u003c/li\u003e\u003cli\u003eIn a pan, combine coconut oil, honey, lemon peel and lemon juice. Bring to boil. Take off the heat when the honey is dissolved. Let cool a little.\u003c/li\u003e\u003cli\u003eIn another bowl, combine oats, flax seeds and chopped nuts. Mix well. Pour the honey mixture over the dry ingredients and stir until the oat mixture is well coated.\u003c/li\u003e\u003cli\u003eSpread the granola evenly on a baking pan, and bake for 20 minutes. Turn over the granola, and stir in the coconut flakes. Bake for another 10 minutes. Let cool and stir in the currants. Keep in airtight container.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":60
  },
  {
    "title":"Candied Spiced Pecans",
    "picture":"https://spoonacular.com/recipeImages/636864-556x370.jpg",
    "ingredients":[
      "allspice",
      "brown sugar",
      "butter",
      "cinnamon",
      "granulated sugar",
      "ground cloves",
      "maple extract",
      "nutmeg",
      "pecans",
      "water"
    ],
    "ingredientInfo":{
      "allspice":{
        "amount":0.25,
        "measuresUnit":"tsp",
        "image":"allspice-ground.jpg",
        "name":"allspice"
      },
      "brown sugar":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"light-brown-sugar.jpg",
        "name":"brown sugar"
      },
      "butter":{
        "amount":0.5,
        "measuresUnit":"stick",
        "image":"butter-sliced.jpg",
        "name":"butter"
      },
      "cinnamon":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"cinnamon.jpg",
        "name":"cinnamon"
      },
      "granulated sugar":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"sugar-in-bowl.png",
        "name":"granulated sugar"
      },
      "ground cloves":{
        "amount":0.25,
        "measuresUnit":"tsp",
        "image":"cloves.jpg",
        "name":"ground cloves"
      },
      "maple extract":{
        "amount":1,
        "measuresUnit":"tsp",
        "image":"vanilla-extract.jpg",
        "name":"maple extract"
      },
      "nutmeg":{
        "amount":0.25,
        "measuresUnit":"tsp",
        "image":"ground-nutmeg.jpg",
        "name":"nutmeg"
      },
      "pecans":{
        "amount":2,
        "measuresUnit":"cups",
        "image":"pecans.jpg",
        "name":"pecans"
      },
      "water":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"water.png",
        "name":"water"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePreheat oven to 300 F.\u003c/li\u003e\u003cli\u003eLine a cookie sheet with tin foil and coat with cooking spray. Set aside.\u003c/li\u003e\u003cli\u003eIn a shallow bowl sift together granulated sugar, cinnamon, allspice, cloves, and nutmeg. Set aside.\u003c/li\u003e\u003cli\u003eIn a small bowl, whisk together maple extract, brown sugar and water.\u003c/li\u003e\u003cli\u003eOver medium heat, melt butter in a skillet. Slowly whisk in brown sugar mixture and bring to a boil. Add pecans and stir to coat. Simmer for about 1-2 minutes until well coated.\u003c/li\u003e\u003cli\u003eWith a slotted spoon, transfer pecans to the reserved sugar and spice bowl. Coat pecans with sugar.\u003c/li\u003e\u003cli\u003ePlace pecans evenly on prepared cookie sheet and bake for 15 minutes. Stir pecans and bake for an additional 15 minutes. Let cool completely before serving.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":30
  },
  {
    "title":"Deviled Eggs With Crab",
    "picture":"https://spoonacular.com/recipeImages/641461-556x370.jpg",
    "ingredients":[
      "celery",
      "dijon mustard",
      "eggs",
      "fresh chives",
      "lemon juice",
      "lump crabmeat",
      "mayonnaise",
      "salt and pepper",
      "sour cream"
    ],
    "ingredientInfo":{
      "celery":{
        "amount":0.5,
        "measuresUnit":"stick",
        "image":"celery.jpg",
        "name":"celery"
      },
      "dijon mustard":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"dijon-mustard.jpg",
        "name":"dijon mustard"
      },
      "eggs":{
        "amount":6,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"eggs"
      },
      "fresh chives":{
        "amount":3,
        "measuresUnit":"tablespoons",
        "image":"fresh-chives.jpg",
        "name":"fresh chives"
      },
      "lemon juice":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"lemon-juice.jpg",
        "name":"lemon juice"
      },
      "lump crabmeat":{
        "amount":4,
        "measuresUnit":"ounces",
        "image":"lump-crabmeat.png",
        "name":"lump crabmeat"
      },
      "mayonnaise":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"mayonnaise.png",
        "name":"mayonnaise"
      },
      "salt and pepper":{
        "amount":6,
        "measuresUnit":"servings",
        "image":"salt-and-pepper.jpg",
        "name":"salt and pepper"
      },
      "sour cream":{
        "amount":2,
        "measuresUnit":"tablespoons",
        "image":"sour-cream.jpg",
        "name":"sour cream"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eIn a medium-sized mixing bowl combine the crabmeat, celery, sour cream, mayonnaise, Dijon mustard, lemon juice, and chives. Stir until well combined.\u003c/li\u003e\u003cli\u003eSeason, to taste, with salt and pepper.\u003c/li\u003e\u003cli\u003eSpoon the crab mixture into the egg halves.\u003c/li\u003e\u003cli\u003eServe immediately or chilled until ready.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":37
  },
  {
    "title":"Zucchini Flutes Piped With Basil Ricotta Mousse",
    "picture":"https://spoonacular.com/recipeImages/665744-556x370.jpg",
    "ingredients":[
      "zucchini",
      "basil leaves",
      "garlic",
      "ricotta",
      "parmesan cheese",
      "extra virgin olive oil"
    ],
    "ingredientInfo":{
      "basil leaves":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"basil.jpg",
        "name":"basil leaves"
      },
      "extra virgin olive oil":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"olive-oil.jpg",
        "name":"extra virgin olive oil"
      },
      "garlic":{
        "amount":2,
        "measuresUnit":"tablespoons",
        "image":"garlic.png",
        "name":"garlic"
      },
      "parmesan cheese":{
        "amount":4,
        "measuresUnit":"tablespoons",
        "image":"parmesan.jpg",
        "name":"parmesan cheese"
      },
      "ricotta":{
        "amount":1.5,
        "measuresUnit":"cups",
        "image":"ricotta.png",
        "name":"ricotta"
      },
      "zucchini":{
        "amount":2,
        "measuresUnit":"medium",
        "image":"zucchini.jpg",
        "name":"zucchini"
      }
    },
    "summary":"Preheat oven to 375 degrees.\nZucchini:  Cut stem tips of zucchinis off and discard. Slice lengthwise into two pieces. Take a teaspoon and hollow out each half... scraping away seeds and core until smooth. Be sure to leave about 1/4 inch of flesh or the flute will be too weak. Set on baking tray.\nMousse:  Put the ricotta, basil, garlic, and half of the shredded Parmesan cheese into a food processor and blend until creamy. If too thick... add just a dab of olive oil. But not too much as you need the mousse to stand firm in the zucchini flutes.\nOnce you reach the desired consistency - scoop the mousse out of processor into a zip lock baggie or pastry bag. If using a baggie, snip off about 1/4\" of the corner and squeeze baggie to pipe out mousse into the hollowed section of the zucchini. Stop short of the end by about 1/2\", as the ricotta will expand when baking. Sprinkle remaining Parmesan cheese along the top of flutes.\nPut tray of mousse-filled flutes on middle rack of oven, baking for 20 minutes at 375 degrees. Remove when zucchini is tender to a fork and the cheese has browned slightly. Once flutes are removed from the oven, sprinkle a few ricotta crumbles across the top and lightly drizzle with extra virgin olive oil. Serve immediately.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":55
  },
  {
    "title":"Lemon square bars",
    "picture":"https://spoonacular.com/recipeImages/649794-556x370.jpg",
    "ingredients":[
      "baking powder",
      "confectioners' sugar",
      "eggs",
      "flour",
      "lemon juice",
      "lemon rind",
      "sugar",
      "unsalted butter",
      "vanilla essence"
    ],
    "ingredientInfo":{
      "baking powder":{
        "amount":0.5,
        "measuresUnit":"teaspoon",
        "image":"white-powder.jpg",
        "name":"baking powder"
      },
      "confectioners' sugar":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"powdered-sugar.jpg",
        "name":"confectioners' sugar"
      },
      "eggs":{
        "amount":2,
        "measuresUnit":"",
        "image":"egg.png",
        "name":"eggs"
      },
      "flour":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"flour.png",
        "name":"flour"
      },
      "lemon juice":{
        "amount":0.75,
        "measuresUnit":"cup",
        "image":"lemon-juice.jpg",
        "name":"lemon juice"
      },
      "lemon rind":{
        "amount":1,
        "measuresUnit":"teaspoon",
        "image":"zest-lemon.jpg",
        "name":"lemon rind"
      },
      "sugar":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      },
      "unsalted butter":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"butter-sliced.jpg",
        "name":"unsalted butter"
      },
      "vanilla essence":{
        "amount":0.5,
        "measuresUnit":"tsp",
        "image":"vanilla-extract.jpg",
        "name":"vanilla essence"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003eIn a medium bowl, stir together 2 cups flour and confectioners' sugar. Blend in the melted butter and vanilla essence.\u003c/li\u003e\u003cli\u003ePress the soft dough with the palm of your hand into the bottom of the greased 9x13 inch pan.\u003c/li\u003e\u003cli\u003eBake the dough in the preheated oven at 170C for 15 minutes, or until golden.\u003c/li\u003e\u003cli\u003eMeanwhile, in a large bowl, beat eggs until light.\u003c/li\u003e\u003cli\u003eCombine the sugar, baking powder and  cup of flour so there will be no flour lumps.\u003c/li\u003e\u003cli\u003eStir the sugar mixture into the eggs.\u003c/li\u003e\u003cli\u003eStir in the lemon juice and lemon rind. Pour over the prepared crust and return to the oven.\u003c/li\u003e\u003cli\u003eBake for an additional 30 minutes or until bars are set.\u003c/li\u003e\u003cli\u003eAllow to cool completely before cutting into bars.\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":7
  },
  {
    "title":"Banana Creme Brulee",
    "picture":"https://spoonacular.com/recipeImages/634070-556x370.jpg",
    "ingredients":[
      "bananas",
      "demerara sugar",
      "egg yolks",
      "heavy cream",
      "sugar",
      "vanilla pod"
    ],
    "ingredientInfo":{
      "bananas":{
        "amount":2,
        "measuresUnit":"",
        "image":"bananas.jpg",
        "name":"bananas"
      },
      "demerara sugar":{
        "amount":1,
        "measuresUnit":"",
        "image":"raw-brown-sugar.png",
        "name":"demerara sugar"
      },
      "egg yolks":{
        "amount":3,
        "measuresUnit":"",
        "image":"egg-yolk.jpg",
        "name":"egg yolks"
      },
      "heavy cream":{
        "amount":450,
        "measuresUnit":"ml",
        "image":"fluid-cream.jpg",
        "name":"heavy cream"
      },
      "sugar":{
        "amount":0.5,
        "measuresUnit":"cup",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      },
      "vanilla pod":{
        "amount":1,
        "measuresUnit":"",
        "image":"vanilla.jpg",
        "name":"vanilla pod"
      }
    },
    "summary":"\u003col\u003e\u003cli\u003ePut the cream and the vanilla in a medium saucepan and bring slowly to the boil.\u003c/li\u003e\u003cli\u003eRemove from the heat just before the cream boils.\u003c/li\u003e\u003cli\u003eSplit the pod in half and scrape out the seeds with the point of a knife.\u003c/li\u003e\u003cli\u003eSlice the bananas thickly and divide between 8 small ramekins.\u003c/li\u003e\u003cli\u003ePut the egg yolks caster sugar and vanilla seeds in a mixing bowl and beat till thick and creamy. Pour the hot milk on to the egg and sugar mixture and stir.\u003c/li\u003e\u003cli\u003eRinse out the milk pan dry and pour in the custard.\u003c/li\u003e\u003cli\u003eHeat stirring slowly and almost constantly until the mixture thickens.\u003c/li\u003e\u003cli\u003eThere are a couple of things to bear in mind: if you make certain that the spoon gets right into the corners of the pan you run less risk of the custard curdling and on no account let the mixture boil otherwise the custard will scramble.\u003c/li\u003e\u003cli\u003ePour the custard through a sieve into the little dishes filling them right to the top.\u003c/li\u003e\u003cli\u003eSet aside to cool then refrigerate overnight.\u003c/li\u003e\u003cli\u003eDust the top of each custard with a thin layer of demerara then place under a very very hot grill for a few seconds until the sugar melts to a shiny caramel.\u003c/li\u003e\u003cli\u003eLeave to cool and harden.\u003c/li\u003e\u003cli\u003eMakes 8\u003c/li\u003e\u003c/ol\u003e",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":18
  },
  {
    "title":"Milky Watermelon Drink",
    "picture":"https://spoonacular.com/recipeImages/716252-556x370.jpg",
    "ingredients":[
      "dates",
      "evaporated milk",
      "ginger",
      "ice",
      "watermelon"
    ],
    "ingredientInfo":{
      "dates":{
        "amount":2,
        "measuresUnit":"pieces",
        "image":"dates.jpg",
        "name":"dates"
      },
      "evaporated milk":{
        "amount":7,
        "measuresUnit":"tablespoons",
        "image":"evaporated-milk.png",
        "name":"evaporated milk"
      },
      "ginger":{
        "amount":0.25,
        "measuresUnit":"teaspoon",
        "image":"ginger.png",
        "name":"ginger"
      },
      "ice":{
        "amount":0.25,
        "measuresUnit":"cup",
        "image":"ice-cubes.png",
        "name":"ice"
      },
      "watermelon":{
        "amount":1,
        "measuresUnit":"cup",
        "image":"watermelon.png",
        "name":"watermelon"
      }
    },
    "summary":"Blend all ingredients together till smooth and serve chilled.P:S вЂ“ If you are wondering where to buy dates, check out supermarkets or the hausa traders at a market. They are more likely to have dates. Does anyone know what dates are called in Hausa? Please leave a comment and help others",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":58
  },
  {
    "title":"Mango Salsa",
    "picture":"https://spoonacular.com/recipeImages/716290-556x370.jpg",
    "ingredients":[
      "cherry tomatoes",
      "green bell pepper",
      "lime juice",
      "mangoes",
      "mint leaves",
      "red bell pepper",
      "red onions",
      "sugar"
    ],
    "ingredientInfo":{
      "cherry tomatoes":{
        "amount":1,
        "measuresUnit":"handful",
        "image":"cherry-tomatoes.png",
        "name":"cherry tomatoes"
      },
      "green bell pepper":{
        "amount":0.5,
        "measuresUnit":"",
        "image":"green-pepper.jpg",
        "name":"green bell pepper"
      },
      "lime juice":{
        "amount":1,
        "measuresUnit":"tablespoon",
        "image":"lime-juice.png",
        "name":"lime juice"
      },
      "mangoes":{
        "amount":2,
        "measuresUnit":"",
        "image":"mango.jpg",
        "name":"mangoes"
      },
      "mint leaves":{
        "amount":3,
        "measuresUnit":"leaves",
        "image":"mint.jpg",
        "name":"mint leaves"
      },
      "red bell pepper":{
        "amount":0.5,
        "measuresUnit":"",
        "image":"red-pepper.jpg",
        "name":"red bell pepper"
      },
      "red onions":{
        "amount":2,
        "measuresUnit":"handfuls",
        "image":"red-onion.png",
        "name":"red onions"
      },
      "sugar":{
        "amount":1,
        "measuresUnit":"pinch",
        "image":"sugar-in-bowl.png",
        "name":"sugar"
      }
    },
    "summary":"Peel and chop your mango into small cubesCube your bell peppers and onions as well and mix in with the mangoes. Add the cherry tomatoes and mix in.sprinkle your sugar and lime juice over it.Chop your scent/mint leaves and mix in and refrigerate.Serve cool as a side dish or if you want something refreshing on a hot day.",
    "readyInMinutes":45,
    "cookingDish":"",
    "cookingMethod":"",
    "rating":91
  }
]`
