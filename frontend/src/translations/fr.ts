export default {
    visibility: {
        public: 'Publique',
        unlisted: 'Non listé',
        private: 'Privé'
    },
    dropzone: {
      text: 'Cliquez ou lachez un fichier ici pour téléverser'
    },
    menu: {
        home: 'Accueil',
        login: 'Connexion',
        register: "S'inscrire",
        server: 'Serveur',
        logout: 'Déconnexion'
    },
    errors: {
        unknown: 'Erreur inconnue',
        passwordNotMatching: 'Les mots de passes ne correspondent pas',
    },
    registration: {
        title: "S'inscrire",
        username: "Pseudo",
        email: "Email",
        password: "Mot de passe",
        repeat_password: "Répéter le mot de passe",
        invite_code: "Code d'invitation",
        button: "S'inscrire",
    },
    login: {
        title: 'Connexion',
        username: 'Pseudo',
        password: 'Mot de passe',
        button: 'Connexion'
    },
    validate: {
        accountValidated: 'Votre compte est validé',
        connect: 'Vous pouvez maintenant vous connecter avec le pseudo'
    },
    my_profile: {
        account: {
            tab_name: 'Mon compte',
            current_password: "Mot de passe actuel",
            new_password: "Nouveau mot de passe",
            update: 'Sauvegarder',
            success: 'Votre compte à été mis à jour avec succès',
        },
        tokens: {
            tab_name: 'Appareils connectés',
            logged_at: 'Connecté à',
            revoked_at: 'Révoqué à',
            last_seen: 'Dernière utilisation',
            never: 'Jamais',
            on: 'le',
        },
        sharex: {
            tab_name: 'ShareX'
        },
    },
    server_settings: {
        invite: {
            tab_name: "Code d'invitation",
            code: "Code d'invitation",
            generate: "Générer",
            list: {
                generated_by: 'Généré par',
                on: 'Le',
            },
            delete_dialog: {
                title: "Supprimer un code d'invitation",
                text: "La suppression de ce code d'invitation empêchera quiconque de s'inscrire avec",
                remove: 'Supprimer',
                cancel: 'Annuler'
            }
        }
    },
    dial: {
        textbin: 'Textbin',
        url_shortener: 'Minimiseur d\'URL'
    },
    browser: {
        folder_editor: {
            create_title: 'Créer un dossier',
            edit_title: 'Modifier un dossier',
            folder_name: 'Nom',
            visibility: 'Visibilité',
            create: 'Créer',
            save: 'Sauvegarder',
            cancel: 'Annuler'
        },
        file_uploader: {
            screen_1: {
                title: 'Téléverser',
                file_name: 'Nom',
                description: 'Description',
                visibility: 'Visibilité',
                cancel: 'Annuler',
                upload: 'Envoyer',
            },
            screen_2: {
                title: 'Succès !',
                text: 'Votre image à été téléversée.',
                scinna_link: 'Lien Scinna',
                raw_link: 'Lien direct',
                close: 'Fermer'
            }
        },
        context: {
            edit: 'Éditer',
            remove: 'Supprimer',
        },
        modals: {
            remove_collection: {
                title: 'Supprimer la collection',
                text: 'Êtes-vous sur de vouloir supprimer cette collection et son contenu ?',
                cancel: 'Annuler',
                delete: 'Supprimer'
            }
        }
    },
    shortener: {
        link: 'Lien',
        send: 'Minimifier !'
    }
};