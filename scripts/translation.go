package groupieTracker

type Translation struct {
	YnovProject         string
	Home                string
	About               string
	Sort                string
	Default             string
	AZ                  string
	CreationDate        string
	NumberOfMembers     string
	Cards               string
	FourCards           string
	ThirteenCards       string
	TwentySixCards      string
	Confirm             string
	Prev                string
	Next                string
	WhatIsThis          string
	WhatIsThisOne       string
	WhoAreYou           string
	WhoAreYouOne        string
	TechnicalStuff      string
	TechnicalStuffOne   string
	TechnicalStuffTwo   string
	TechnicalStuffThree string
	WhoDidWhat          string
	WhoDidWhatOne       string
	WhoDidWhatTwo       string
	WhoDidWhatThree     string
	WhoDidWhatFour      string
	Copyright           string
	CopyrightOne        string
	History             string
	HistoryOne          string
	HistoryTwo          string
	Members             string
	Error               string
	Artist              string
	NotFound            string
}

// Start of translation function
func TranslationMode(language string) Translation {

	var translatedData Translation

	if language == "english" {
		translatedData.YnovProject = "Ynov Project"
		translatedData.Home = "Home"
		translatedData.About = "About"
		translatedData.Sort = "Sort"
		translatedData.Default = "Default"
		translatedData.AZ = "A-Z"
		translatedData.CreationDate = "Creation date"
		translatedData.NumberOfMembers = "Number of members"
		translatedData.Cards = "Cards"
		translatedData.FourCards = "4 cards"
		translatedData.ThirteenCards = "13 cards"
		translatedData.TwentySixCards = "26 cards"
		translatedData.Confirm = "Confirm"
		translatedData.Prev = "Prev"
		translatedData.Next = "Next"
		translatedData.WhatIsThis = "What is this ?"
		translatedData.WhatIsThisOne = "Groupie-Tracker is a student project based on an API. You will have access to information about dozens of different artists, albums, members, dates, etc."
		translatedData.WhoAreYou = "Who are you ?"
		translatedData.WhoAreYouOne = "We are students in computer science at Ynov Nantes. Members :"
		translatedData.TechnicalStuff = "Technical stuff ?"
		translatedData.TechnicalStuffOne = "The whole project is coded in"
		translatedData.TechnicalStuffTwo = "on the backend. We use a"
		translatedData.TechnicalStuffThree = "specially created for our student project. And we work in groups with"
		translatedData.WhoDidWhat = "Who did what ?"
		translatedData.WhoDidWhatOne = ": connection api artist/date/location/relation, search system by name with error management, sorting of dates, display of data on the html pages."
		translatedData.WhoDidWhatTwo = ": HTML/CSS, filtering by date of creation/number of artists, filtering by alphabetical order, display of data and responsiveness, golang server, html/server link."
		translatedData.WhoDidWhatThree = ": HTML/CSS, golang server, responsiveness, pagination by number of element 8/16/32, display of data on the html page, relation api connection, (official)README."
		translatedData.WhoDidWhatFour = ": (draft)README, reflection on pagination, artist api connection (not implemented on the project)."
		translatedData.Copyright = "Copyright ?"
		translatedData.CopyrightOne = "The data and images in the API are used without claim of ownership and belong to their respective owners."
		translatedData.History = "History"
		translatedData.HistoryOne = "was create in"
		translatedData.HistoryTwo = "The first album was relise in"
		translatedData.Members = "Members"
		translatedData.Error = "Error"
		translatedData.Artist = "Artist"
		translatedData.NotFound = "not found..."

		return translatedData

	} else if language == "french" {

		translatedData.YnovProject = "Projet Ynov"
		translatedData.Home = "Accueil"
		translatedData.About = "À propos"
		translatedData.Sort = "Trier"
		translatedData.Default = "Par défaut"
		translatedData.AZ = "A-Z"
		translatedData.CreationDate = "Date de création"
		translatedData.NumberOfMembers = "Effectif"
		translatedData.Cards = "Cartes"
		translatedData.FourCards = "4 cartes"
		translatedData.ThirteenCards = "13 cartes"
		translatedData.TwentySixCards = "26 cartes"
		translatedData.Confirm = "Confirmer"
		translatedData.Prev = "Préc"
		translatedData.Next = "Suiv"
		translatedData.WhatIsThis = "Qu'est-ce que c'est ?"
		translatedData.WhatIsThisOne = "Groupie-Tracker est un projet étudiant basé sur une API. Vous aurez accès à des informations sur des dizaines d'artistes différents, des albums, des membres, des dates, etc."
		translatedData.WhoAreYou = "Qui êtes-vous ?"
		translatedData.WhoAreYouOne = "Nous sommes étudiants en informatique à Ynov Nantes. Membres :"
		translatedData.TechnicalStuff = "Questions techniques ?"
		translatedData.TechnicalStuffOne = "L'ensemble du projet est codé en"
		translatedData.TechnicalStuffTwo = "pour le backend. Nous utilisons un"
		translatedData.TechnicalStuffThree = "spécialement créé pour notre projet étudiant. Et nous travaillons en groupes avec"
		translatedData.WhoDidWhat = "Qui a fait quoi ?"
		translatedData.WhoDidWhatOne = ": connexion api artiste/date/localisation/relation, système de recherche par nom avec gestion des erreurs, tri des dates, affichage des données sur les pages html."
		translatedData.WhoDidWhatTwo = ": HTML/CSS, filtrage par date de création/nombre d'artistes, filtrage par ordre alphabétique, affichage des données et réactivité, serveur golang, lien html/serveur."
		translatedData.WhoDidWhatThree = ": HTML/CSS, serveur golang, réactivité, pagination par nombre d'élément 8/16/32, affichage des données sur la page html, connexion api relation, (officiel)README."
		translatedData.WhoDidWhatFour = ": (brouillon)README, réflexion sur la pagination, connexion à l'api de l'artiste (non implémenté sur le projet)."
		translatedData.Copyright = "Droits d'auteur ?"
		translatedData.CopyrightOne = "Les données et les images contenues dans l'API sont utilisées sans revendication de propriété et appartiennent à leurs propriétaires respectifs."
		translatedData.History = "Histoire"
		translatedData.HistoryOne = "a été créée en"
		translatedData.HistoryTwo = "Le premier album a été publié en"
		translatedData.Members = "Membres"
		translatedData.Error = "Erreur"
		translatedData.Artist = "Artiste"
		translatedData.NotFound = "non trouvé..."

		return translatedData
	}

	return translatedData
}
