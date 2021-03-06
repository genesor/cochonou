package bolt_test

import (
	"testing"

	"github.com/asdine/storm"
	"github.com/stretchr/testify/require"

	"github.com/genesor/cochonou"
	"github.com/genesor/cochonou/bolt"
)

func setup(t *testing.T, db storm.Node) (*bolt.RedirectionStore, func()) {
	tx, err := db.Begin(true)
	if err != nil {
		t.Fatal(err)
	}
	store := &bolt.RedirectionStore{DB: tx}

	return store, func() { tx.Rollback() }
}

func TestSave(t *testing.T) {
	db, err := storm.Open("../cochonou_test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("OK - Save", func(t *testing.T) {
		store, rollback := setup(t, db)
		defer rollback()

		redir := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		redir2 := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon2",
		}

		err := store.Save(redir)
		require.NoError(t, err)
		require.NotEqual(t, 0, redir.ID)

		err = store.Save(redir2)
		require.NoError(t, err)
		require.NotEqual(t, 0, redir2.ID)
		require.NotEqual(t, redir.ID, redir2.ID)
	})

	t.Run("NOK - Duplicate", func(t *testing.T) {
		store, rollback := setup(t, db)
		defer rollback()

		redir := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		redir2 := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		err := store.Save(redir)
		require.NoError(t, err)
		require.NotEqual(t, 0, redir.ID)

		err = store.Save(redir2)
		require.NotNil(t, err)
		require.Equal(t, cochonou.ErrSubDomainUsed, err)
		require.Equal(t, 0, redir2.ID)
	})
}

func TestGetBySubDomain(t *testing.T) {
	db, err := storm.Open("../cochonou_test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("OK - Exists", func(t *testing.T) {
		store, rollback := setup(t, db)
		defer rollback()

		redir := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		err := store.Save(redir)
		require.NoError(t, err)
		require.NotEqual(t, 0, redir.ID)

		redir2, err := store.GetBySubDomain("cochon")
		require.NoError(t, err)
		require.Equal(t, "cochon", redir2.SubDomain)
		require.Equal(t, "http://sadoma.so/", redir2.URL)
	})

	t.Run("NOK - Doesn't exists", func(t *testing.T) {
		store, rollback := setup(t, db)
		defer rollback()

		redir2, err := store.GetBySubDomain("cochon")
		require.Error(t, err)
		require.Equal(t, cochonou.ErrNotFound, err)
		require.Nil(t, redir2)
	})
}

func TestGetAll(t *testing.T) {
	db, err := storm.Open("../cochonou_test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("OK", func(t *testing.T) {
		store, rollback := setup(t, db)
		defer rollback()

		redir := &cochonou.Redirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		err := store.Save(redir)
		require.NoError(t, err)
		require.NotEqual(t, 0, redir.ID)

		list, err := store.GetAll()
		require.NoError(t, err)
		require.Len(t, list, 1)
		require.Equal(t, "cochon", list[0].SubDomain)
		require.Equal(t, "http://sadoma.so/", list[0].URL)
	})
}
