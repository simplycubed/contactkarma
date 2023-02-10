/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/simplycubed/contactkarma/contacts/adapters"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Adds test data to firestore",
	Long:  `Adds test data to firestore`,
	Run: func(cmd *cobra.Command, args []string) {
		userId, err := cmd.Flags().GetString("user-id")
		if err != nil {
			log.Fatal(err)
		}
		if userId == "" {
			userId = gofakeit.UUID()
		}
		err = Seed(userId)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Seed(userId string) (err error) {
	dbpool := adapters.NewFirestore()
	userRepo := firestore.NewUserFirestore(dbpool)
	contactRepo := firestore.NewDefaultContactsFirestore(dbpool)
	noteRepo := firestore.NewNoteFirestore(dbpool)
	tagRepo := firestore.NewTagFirestore(dbpool)
	contactSourceRepo := firestore.NewContactSourceFirestore(dbpool)
	unifiedRepo := firestore.NewUnifiedContactFirestore(dbpool)
	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedRepo, linkSuggestionRepo)
	unifiedContactService := application.NewUnifiedContactService(unifiedRepo, linkSuggestionService, contactLogRepo)
	seeder := &Seeder{
		contactRepo:           contactRepo,
		userRepo:              userRepo,
		tagRepo:               tagRepo,
		contactSourceRepo:     contactSourceRepo,
		noteRepo:              noteRepo,
		unifiedContactService: unifiedContactService,
	}
	err = seeder.AddUserData(userId)
	if err != nil {
		return
	}
	log.Println("User generated successfully")
	return
}

type Seeder struct {
	contactRepo           *firestore.DefaultContactsFirestore
	userRepo              *firestore.UserFirestore
	noteRepo              *firestore.NoteFirestore
	tagRepo               *firestore.TagFirestore
	contactSourceRepo     *firestore.ContactSourceFirestore
	unifiedContactService *application.UnifiedContactService
}

// adds a user and contacts
func (s *Seeder) AddUserData(id string) (err error) {
	user, err := s.AddUser(id)
	if err != nil {
		return
	}

	contactsCount := 1 + rand.Intn(10)
	err = s.AddContacts(contactsCount, user.ID)
	return
}

func (s *Seeder) AddContacts(count int, userId domain.UserID) (err error) {
	for i := 0; i < count; i++ {
		var contact *domain.Contact
		contact, err = s.AddContact(userId)
		if err != nil {
			return
		}
		var unifiedContact *domain.Unified
		unifiedContact, err = s.unifiedContactService.SyncContactToUnified(context.Background(), userId, domain.Google, domain.ContactSourceID(domain.Default), domain.ContactID(contact.ID), *contact)
		if err != nil {
			return
		}
		// add notes
		notesCount := 1 + rand.Intn(10)
		err = s.AddNotes(notesCount, userId, unifiedContact.ID)
		if err != nil {
			return
		}
		// add tags
		tagsCount := 1 + rand.Intn(10)
		err = s.AddTags(tagsCount, userId, unifiedContact.ID)
		if err != nil {
			return
		}
	}
	return
}

func (s *Seeder) AddNotes(count int, userId domain.UserID, contactId domain.UnifiedId) (err error) {
	for i := 0; i < count; i++ {
		_, err = s.AddNote(userId, contactId)
		if err != nil {
			return
		}
	}
	return
}

func (s *Seeder) AddTags(count int, userId domain.UserID, contactId domain.UnifiedId) (err error) {
	for i := 0; i < count; i++ {
		_, err = s.AddTag(userId, contactId)
		if err != nil {
			return
		}
	}
	return
}

func (s *Seeder) AddUser(id string) (user *domain.User, err error) {
	// create a user
	fakeUser := domain.User{}
	err = gofakeit.Struct(&fakeUser)
	if err != nil {
		return
	}
	user, err = s.userRepo.SaveUser(context.Background(), domain.UserID(id), fakeUser)
	return
}

func (s *Seeder) AddContact(userId domain.UserID) (contact *domain.Contact, err error) {
	fakeUser := domain.Contact{}
	err = gofakeit.Struct(&fakeUser)
	if err != nil {
		return
	}
	contact, err = s.contactRepo.SaveContact(context.Background(), userId, fakeUser)
	return
}

func (s *Seeder) AddNote(userId domain.UserID, contactId domain.UnifiedId) (note *domain.Note, err error) {
	fakeNote := domain.Note{}
	err = gofakeit.Struct(&fakeNote)
	if err != nil {
		return
	}
	note, err = s.noteRepo.SaveNote(context.Background(), userId, contactId, fakeNote)
	return
}

func (s *Seeder) AddTag(userId domain.UserID, contactId domain.UnifiedId) (tag *domain.Tag, err error) {
	fakeTag := domain.Tag{}
	err = gofakeit.Struct(&fakeTag)
	if err != nil {
		return
	}
	tag, err = s.tagRepo.SaveTag(context.Background(), userId, contactId, fakeTag)
	return
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	seedCmd.Flags().String("user-id", "", "set id of user being created")
}
