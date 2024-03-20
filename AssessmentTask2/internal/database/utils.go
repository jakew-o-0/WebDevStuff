package database

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var dbSchema string

func CreateFromSchema(db *sql.DB) error {
    _,err := db.ExecContext(context.Background(), dbSchema)
    return err
}

type ZooAttraction struct {
    Title   string
    Content string
}

func AttractionDummyData() []ZooAttraction {
    return []ZooAttraction{
        {
            Title:   "Lion's Den",
            Content: "Come face to face with the king of the jungle in our state-of-the-art lion habitat. Watch them roam and play in their naturalistic environment.",
        },
        {
            Title:   "Monkey Island",
            Content: "Swing by Monkey Island and observe our troop of playful primates. From agile gibbons to mischievous capuchins, there's always something entertaining happening here.",
        },
        {
            Title:   "Aquatic Wonders",
            Content: "Dive into the fascinating world of aquatic life in our expansive aquarium. Marvel at colorful tropical fish, graceful rays, and even catch a glimpse of our majestic sharks.",
        },
        // Add more attractions here
        {
            Title:   "Birds of Paradise",
            Content: "Visit our aviary and witness the beauty of exotic birds from around the globe. From vibrant macaws to majestic eagles, you'll be mesmerized by their colors and grace.",
        },
        {
            Title:   "Reptile House",
            Content: "Explore the mysterious world of reptiles in our reptile house. From slithering snakes to ancient turtles, learn about these fascinating creatures and their importance to our ecosystem.",
        },
        {
            Title:   "Butterfly Garden",
            Content: "Step into our butterfly garden and be surrounded by a kaleidoscope of fluttering colors. Witness the delicate beauty of these winged wonders as they flit among the flowers.",
        },
        {
            Title:   "African Safari",
            Content: "Embark on an African safari adventure and encounter giraffes, zebras, and rhinos in our expansive savannah exhibit. Experience the thrill of the Serengeti right here at our zoo.",
        },
        {
            Title:   "Penguin Parade",
            Content: "Join our playful penguins as they waddle and slide in their icy habitat. Watch them dive into the water and perform acrobatic feats that will leave you in awe.",
        },
        {
            Title:   "Elephant Encounters",
            Content: "Meet our gentle giants up close in our elephant enclosure. Learn about their intelligence, social behaviors, and the conservation efforts aimed at protecting these magnificent creatures.",
        },
        {
            Title:   "Tropical Rainforest",
            Content: "Immerse yourself in the sights and sounds of the tropical rainforest in our lush exhibit. Encounter colorful frogs, majestic jaguars, and towering trees teeming with life.",
        },
        {
            Title:   "Savannah Safari",
            Content: "Embark on a safari through the African savannah and witness lions, giraffes, and zebras in their natural habitat. Experience the beauty and diversity of the African plains right here at our zoo.",
        },
        {
            Title:   "Koala Kingdom",
            Content: "Enter the enchanting world of koalas in our Koala Kingdom exhibit. Watch these adorable marsupials lounge in eucalyptus trees and learn about their unique lifestyles and conservation status.",
        },
        {
            Title:   "Gorilla Grove",
            Content: "Venture into Gorilla Grove and observe our family of Western lowland gorillas. Learn about these intelligent primates and the efforts being made to protect them in the wild.",
        },
        {
            Title:   "Arctic Adventure",
            Content: "Explore the frozen landscapes of the Arctic in our immersive exhibit. Encounter polar bears, seals, and arctic foxes as you learn about the challenges facing these animals in a changing climate.",
        },
        {
            Title:   "Australian Outback",
            Content: "Journey to the Australian Outback and discover unique wildlife such as kangaroos, emus, and wombats. Learn about the diverse ecosystems of Australia and the conservation efforts underway to protect them.",
        },
        {
            Title:   "Amazon Rainforest",
            Content: "Dive into the depths of the Amazon Rainforest and encounter an array of exotic creatures, from colorful macaws to elusive jaguars. Learn about the importance of preserving this vital ecosystem.",
        },
        {
            Title:   "Panda Paradise",
            Content: "Visit our Panda Paradise and marvel at the adorable antics of our resident giant pandas. Learn about conservation efforts to protect these beloved bears and their bamboo forest habitat.",
        },
        {
            Title:   "Sloth Sanctuary",
            Content: "Slow down and relax in our Sloth Sanctuary, where you can observe these gentle creatures as they move at their own leisurely pace. Learn about their unique adaptations and the importance of preserving their rainforest home.",
        },
        {
            Title:   "Asian Trek",
            Content: "Embark on an Asian adventure and encounter majestic tigers, playful red pandas, and elusive snow leopards. Learn about the rich biodiversity of Asia and the conservation efforts underway to protect it.",
        },
        {
            Title:   "Wildlife Carousel",
            Content: "Take a spin on our Wildlife Carousel and choose from a variety of beautifully crafted animal figures. Fun for all ages, this classic ride is a favorite among visitors.",
        },
    }
}


type ZooFacility struct {
    Title   string
    Content string
}

func FacilityDummyData() []ZooFacility {
    return []ZooFacility{
        {
            Title:   "Visitor Center",
            Content: "Start your zoo adventure at our visitor center. Pick up maps, learn about our conservation efforts, and get information about upcoming events and attractions.",
        },
        {
            Title:   "Restrooms",
            Content: "Conveniently located throughout the zoo, our restrooms are clean and accessible. Look for signage or ask staff for directions to the nearest facilities.",
        },
        {
            Title:   "Gift Shop",
            Content: "Take home a souvenir from your zoo visit at our gift shop. Browse a wide selection of plush toys, apparel, and eco-friendly gifts for all ages.",
        },
        // Add more facilities here
        {
            Title:   "Food Court",
            Content: "Refuel and recharge at our food court. Choose from a variety of dining options, including snacks, sandwiches, and refreshing beverages.",
        },
        {
            Title:   "First Aid Station",
            Content: "In case of emergencies, our first aid station is staffed with trained medical personnel. Look for the red cross symbol or ask any staff member for assistance.",
        },
        {
            Title:   "Playground",
            Content: "Let the kids burn off some energy at our playground. Featuring slides, swings, and climbing structures, it's the perfect spot for young adventurers to take a break.",
        },
        {
            Title:   "Picnic Area",
            Content: "Enjoy a leisurely picnic in our designated picnic area. Bring your own lunch or purchase snacks from nearby vendors and relax amid the natural beauty of the zoo.",
        },
        {
            Title:   "Parking Lot",
            Content: "Park your car in our spacious parking lot conveniently located near the zoo entrance. Parking fees may apply, so be sure to check signage for details.",
        },
        {
            Title:   "Information Kiosks",
            Content: "Find your way around the zoo with ease using our information kiosks. Interactive maps and touch-screen displays provide up-to-date information on attractions and amenities.",
        },
        {
            Title:   "Animal Encounter Zone",
            Content: "Get up close and personal with our animal ambassadors in the Animal Encounter Zone. Meet reptiles, birds, and small mammals while learning about wildlife conservation.",
        },
        {
            Title:   "Education Center",
            Content: "Expand your knowledge of wildlife conservation and ecology at our education center. Join guided tours, workshops, and interactive exhibits led by our expert staff.",
        },
        {
            Title:   "Event Venue",
            Content: "Host your next event at our zoo's versatile event venue. Whether it's a birthday party, corporate gathering, or wedding celebration, our facilities can accommodate your needs.",
        },
        {
            Title:   "Tram Service",
            Content: "Hop aboard our tram service for a convenient way to explore the zoo. Sit back, relax, and enjoy a narrated tour as you travel between exhibits and attractions.",
        },
        {
            Title:   "Accessibility Services",
            Content: "We're committed to making the zoo accessible to all visitors. Ask about our accessibility services, including wheelchair rentals and assistance for guests with special needs.",
        },
        {
            Title:   "Lost and Found",
            Content: "Misplaced something during your visit? Visit our lost and found to reclaim your lost items. Our staff will do their best to reunite you with your belongings.",
        },
        {
            Title:   "Pet Kennel",
            Content: "Planning to visit the zoo but can't bring your pet along? Leave them in the care of our pet kennel, where they'll receive food, water, and attention while you explore.",
        },
        {
            Title:   "Train Station",
            Content: "Embark on a scenic journey around the zoo aboard our train. Sit back and relax as you enjoy panoramic views of animal habitats and lush landscapes.",
        },
        {
            Title:   "Photography Booth",
            Content: "Capture memories of your zoo visit at our photography booth. Professional photographers will snap your picture with your favorite animal, creating lasting keepsakes.",
        },
        {
            Title:   "ATM",
            Content: "Need cash for souvenirs or snacks? Visit our ATM conveniently located near the zoo entrance. Withdrawals are available in multiple denominations.",
        },
        {
            Title:   "Bicycle Rental",
            Content: "Explore the zoo at your own pace with our bicycle rental service. Rent a bike for the day and pedal your way through our scenic trails and pathways.",
        },
    }
}


type AnimalFactFile struct {
    Name        string
    Description string
    Habitat     string
    Diet        string
    ConservationStatus string
}

func CreateDummyFactFiles() []AnimalFactFile {
    return []AnimalFactFile{
        {
            Name:        "African Elephant",
            Description: "The African elephant is the largest land animal on Earth. Known for its long trunk, large ears, and tusks, this gentle giant plays a crucial role in its ecosystem.",
            Habitat:     "African savannahs, forests, and grasslands",
            Diet:        "Herbivore - primarily eats grasses, leaves, and fruits",
            ConservationStatus: "Vulnerable",
        },
        {
            Name:        "Bengal Tiger",
            Description: "The Bengal tiger is the largest of the tiger subspecies and is known for its distinctive orange coat with black stripes. A powerful predator, it is capable of taking down large prey.",
            Habitat:     "Subtropical and tropical forests",
            Diet:        "Carnivore - hunts deer, wild pigs, and other large mammals",
            ConservationStatus: "Endangered",
        },
        {
            Name:        "Giant Panda",
            Description: "The giant panda is a beloved symbol of conservation efforts worldwide. With its distinctive black and white fur, it spends much of its time eating bamboo.",
            Habitat:     "Bamboo forests in central China",
            Diet:        "Mainly herbivore - feeds almost exclusively on bamboo",
            ConservationStatus: "Vulnerable",
        },
        // Add more animal fact files here
        {
            Name:        "African Lion",
            Description: "The African lion is a majestic predator known for its strength and agility. Lions live in social groups called prides and hunt cooperatively to bring down large prey.",
            Habitat:     "Grasslands and savannahs of sub-Saharan Africa",
            Diet:        "Carnivore - primarily hunts large ungulates such as zebras and antelopes",
            ConservationStatus: "Vulnerable",
        },
        {
            Name:        "Giraffe",
            Description: "The giraffe is the tallest land animal, with its long neck and legs allowing it to reach leaves high up in trees. These gentle giants are known for their distinctive spotted coat.",
            Habitat:     "Open woodlands and savannahs of Africa",
            Diet:        "Herbivore - feeds on leaves, shoots, and fruits from trees and shrubs",
            ConservationStatus: "Vulnerable",
        },
        {
            Name:        "Red Panda",
            Description: "The red panda is a small arboreal mammal native to the eastern Himalayas and southwestern China. Despite its name, it is not closely related to the giant panda.",
            Habitat:     "Temperate forests with bamboo understories",
            Diet:        "Mainly herbivore - feeds primarily on bamboo, as well as berries, fruits, and insects",
            ConservationStatus: "Endangered",
        },
        {
            Name:        "Gorilla",
            Description: "Gorillas are the largest living primates and are known for their strength and intelligence. They live in close-knit family groups led by a dominant silverback male.",
            Habitat:     "Forests and jungles of central Africa",
            Diet:        "Herbivore - primarily eats leaves, stems, fruits, and occasionally insects",
            ConservationStatus: "Endangered",
        },
        {
            Name:        "Orangutan",
            Description: "Orangutans are large arboreal apes native to the rainforests of Borneo and Sumatra. They are highly intelligent and are known for their distinctive reddish-brown fur.",
            Habitat:     "Tropical rainforests",
            Diet:        "Mainly frugivore - feeds primarily on fruits, as well as leaves, bark, and insects",
            ConservationStatus: "Critically Endangered",
        },
        {
            Name:        "Polar Bear",
            Description: "The polar bear is a powerful predator native to the Arctic region. With its thick fur and layer of blubber, it is well-adapted to its icy environment.",
            Habitat:     "Arctic sea ice and surrounding coastal areas",
            Diet:        "Carnivore - primarily hunts seals and other marine mammals",
            ConservationStatus: "Vulnerable",
        },
        {
            Name:        "Sloth",
            Description: "Sloths are slow-moving arboreal mammals native to the rainforests of Central and South America. They spend most of their lives hanging upside down from trees.",
            Habitat:     "Tropical rainforests",
            Diet:        "Herbivore - feeds primarily on leaves, shoots, and fruits",
            ConservationStatus: "Least Concern",
        },
    }
}

