package bolt_test

import (
	"testing"

	"github.com/asdine/storm"

	"github.com/genesor/cochonou"
	"github.com/genesor/cochonou/bolt"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	db, err := storm.Open("../cochonou_test.db")
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	t.Run("OK - Save", func(t *testing.T) {
		store := bolt.ImageRedirectionStore{DB: tx}

		redir := &cochonou.ImageRedirection{
			URL:       "http://sadoma.so/",
			SubDomain: "cochon",
		}

		redir2 := &cochonou.ImageRedirection{
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
}
