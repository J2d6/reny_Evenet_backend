-- ================================================
-- TABLE : type_evenement
-- ================================================
CREATE TABLE type_evenement (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nom VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);

-- ================================================
-- TABLE : lieu
-- ================================================
CREATE TABLE lieu (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nom VARCHAR(150) NOT NULL,
    adresse TEXT NOT NULL,
    ville VARCHAR(100) NOT NULL,
    capacite INT CHECK (capacite >= 0 OR capacite IS NULL)
);

-- ================================================
-- TABLE : evenement
-- ================================================
CREATE TABLE evenement (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    titre VARCHAR(150) NOT NULL,
    description TEXT,
    date_debut TIMESTAMP NOT NULL,
    date_fin TIMESTAMP NOT NULL,
    type_id UUID NOT NULL,
    lieu_id UUID NOT NULL,
    CONSTRAINT fk_evenement_type
        FOREIGN KEY (type_id) REFERENCES type_evenement(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_evenement_lieu
        FOREIGN KEY (lieu_id) REFERENCES lieu(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_dates
        CHECK (date_fin >= date_debut)
);

-- ================================================
-- TABLE : type_place
-- ================================================
CREATE TABLE type_place (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nom VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    avantages TEXT
);

-- ================================================
-- TABLE : etat_place
-- ================================================
CREATE TABLE etat_place (
    code VARCHAR(20) PRIMARY KEY,
    description TEXT NOT NULL
);

-- ================================================
-- TABLE : tarif
-- ================================================
CREATE TABLE tarif (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prix DECIMAL(10,2) NOT NULL CHECK (prix >= 0),
    nombre_places INT NOT NULL CHECK (nombre_places >= 0),
    evenement_id UUID NOT NULL,
    type_place_id UUID NOT NULL,
    CONSTRAINT fk_tarif_evenement
        FOREIGN KEY (evenement_id) REFERENCES evenement(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_tarif_type_place
        FOREIGN KEY (type_place_id) REFERENCES type_place(id)
        ON DELETE CASCADE,
    CONSTRAINT uq_tarif UNIQUE (evenement_id, type_place_id)
);

-- ================================================
-- TABLE : place
-- ================================================
CREATE TABLE place (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    numero VARCHAR(100) NOT NULL,
    etat_code VARCHAR(20) NOT NULL DEFAULT 'disponible',
    tarif_id UUID NOT NULL,
    CONSTRAINT fk_place_tarif
        FOREIGN KEY (tarif_id) REFERENCES tarif(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_place_etat
        FOREIGN KEY (etat_code) REFERENCES etat_place(code)
        ON DELETE RESTRICT,
    CONSTRAINT uq_place_numero_tarif UNIQUE (tarif_id, numero)
);

-- ================================================
-- TABLE : audit_place (pour le suivi des changements)
-- ================================================
CREATE TABLE audit_place (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    place_id UUID NOT NULL,
    ancien_etat VARCHAR(20),
    nouvel_etat VARCHAR(20) NOT NULL,
    date_changement TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    utilisateur VARCHAR(100),
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    FOREIGN KEY (ancien_etat) REFERENCES etat_place(code),
    FOREIGN KEY (nouvel_etat) REFERENCES etat_place(code)
);

-- ================================================
-- TABLE : fichier_evenement (avec stockage BYTEA)
-- ================================================
CREATE TABLE fichier_evenement (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evenement_id UUID NOT NULL,
    nom_fichier VARCHAR(255) NOT NULL,
    type_mime VARCHAR(100) NOT NULL,
    taille_bytes BIGINT NOT NULL CHECK (taille_bytes > 0),
    type_fichier VARCHAR(50) NOT NULL CHECK (type_fichier IN ('photo', 'affiche', 'document')),
    donnees_binaire BYTEA NOT NULL,
    date_upload TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_fichier_evenement
        FOREIGN KEY (evenement_id) REFERENCES evenement(id)
        ON DELETE CASCADE
);


-- Insertion des états de place prédéfinis
INSERT INTO etat_place (code, description) VALUES
('disponible', 'Place disponible à la vente'),
('reservee', 'Place réservée temporairement'),
('vendue', 'Place vendue'),
('annulee', 'Place annulée/invalide'),
('maintenance', 'Place en maintenance');

-- Insertion des types d'événement
INSERT INTO type_evenement (nom, description) VALUES
('Concert', 'Événement musical avec artistes sur scène'),
('Conference', 'Événement de présentation et d''échanges'),
('Spectacle', 'Représentation théâtrale ou artistique'),
('Foire', 'Événement commercial avec exposants'),
('Seminaire', 'Séminaire professionnel'),
('Exposition', 'Présentation d''œuvres ou de produits');

-- Insertion des types de place
INSERT INTO type_place (nom, description, avantages) VALUES
('VIP', 'Place premium avec avantages exclusifs', 'Accès lounge, parking dédié, restauration incluse'),
('Standard', 'Place classique standard', 'Accès à l''événement, siège standard'),
('Économique', 'Place à prix réduit', 'Accès basique à l''événement'),
('Premium', 'Place premium confort', 'Sièges plus larges, service prioritaire');



-- ================================================
-- INDEX pour les performances
-- ================================================
CREATE INDEX idx_evenement_dates ON evenement(date_debut, date_fin);
CREATE INDEX idx_evenement_lieu ON evenement(lieu_id);
CREATE INDEX idx_evenement_type ON evenement(type_id);
CREATE INDEX idx_place_etat ON place(etat_code);
CREATE INDEX idx_tarif_evenement ON tarif(evenement_id);
CREATE INDEX idx_tarif_type_place ON tarif(type_place_id);
CREATE INDEX idx_place_tarif ON place(tarif_id);
CREATE INDEX idx_place_numero ON place(numero);
CREATE INDEX idx_evenement_titre ON evenement(titre);
CREATE INDEX idx_audit_place_id ON audit_place(place_id);
CREATE INDEX idx_audit_date ON audit_place(date_changement);
CREATE INDEX idx_fichier_evenement_id ON fichier_evenement(evenement_id);
CREATE INDEX idx_fichier_type ON fichier_evenement(type_fichier);




-- ================================================
-- FONCTION : Validation des paramètres de création
-- ================================================
CREATE OR REPLACE FUNCTION valider_parametres_creation(
    p_titre VARCHAR(150),
    p_description TEXT,
    p_date_debut TIMESTAMP,
    p_date_fin TIMESTAMP,
    p_type_id UUID,
    p_tarifs JSONB
) RETURNS VOID AS $$
BEGIN
    -- Validation du titre
    IF p_titre IS NULL OR trim(p_titre) = '' THEN
        RAISE EXCEPTION 'Le titre est obligatoire';
    END IF;
    
    IF length(trim(p_titre)) > 150 THEN
        RAISE EXCEPTION 'Le titre ne peut pas dépasser 150 caractères';
    END IF;
    
    -- Validation des dates
    IF p_date_debut IS NULL OR p_date_fin IS NULL THEN
        RAISE EXCEPTION 'Les dates de début et fin sont obligatoires';
    END IF;
    
    IF p_date_debut >= p_date_fin THEN
        RAISE EXCEPTION 'La date de début doit être avant la date de fin';
    END IF;
    
    -- Validation du type d'événement
    IF p_type_id IS NULL OR NOT EXISTS (
        SELECT 1 FROM type_evenement WHERE id = p_type_id
    ) THEN
        RAISE EXCEPTION 'Le type d''événement est invalide';
    END IF;
    
    -- Validation des tarifs
    IF p_tarifs IS NULL OR jsonb_array_length(p_tarifs) = 0 THEN
        RAISE EXCEPTION 'Au moins un tarif doit être spécifié';
    END IF;
    
    -- Validation de la structure JSON des tarifs
    IF NOT EXISTS (
        SELECT 1 FROM jsonb_array_elements(p_tarifs) AS tarif
        WHERE tarif ? 'type_place_id' 
          AND tarif ? 'prix' 
          AND tarif ? 'nombre_places'
    ) THEN
        RAISE EXCEPTION 'Structure JSON des tarifs invalide. Champs requis: type_place_id, prix, nombre_places';
    END IF;
    
    -- Validation détaillée des tarifs
    PERFORM valider_tarifs(p_tarifs);
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Calcul du total des places demandées
-- ================================================
CREATE OR REPLACE FUNCTION calculer_total_places(p_tarifs JSONB)
RETURNS INTEGER AS $$
DECLARE
    total INTEGER := 0;
    tarif_record JSONB;
BEGIN
    FOR tarif_record IN SELECT * FROM jsonb_array_elements(p_tarifs) 
    LOOP
        total := total + COALESCE((tarif_record->>'nombre_places')::INTEGER, 0);
    END LOOP;
    
    RETURN total;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Validation détaillée des tarifs
-- ================================================
CREATE OR REPLACE FUNCTION valider_tarifs(p_tarifs JSONB)
RETURNS VOID AS $$
DECLARE
    tarif_record JSONB;
    type_place_id UUID;
    prix DECIMAL(10,2);
    nombre_places INTEGER;
BEGIN
    FOR tarif_record IN SELECT * FROM jsonb_array_elements(p_tarifs) 
    LOOP
        -- Extraction des valeurs
        type_place_id := (tarif_record->>'type_place_id')::UUID;
        prix := (tarif_record->>'prix')::DECIMAL;
        nombre_places := (tarif_record->>'nombre_places')::INTEGER;
        
        -- Validation du type de place
        IF type_place_id IS NULL OR NOT EXISTS (
            SELECT 1 FROM type_place WHERE id = type_place_id
        ) THEN
            RAISE EXCEPTION 'Type de place invalide: %', type_place_id;
        END IF;
        
        -- Validation du prix
        IF prix IS NULL OR prix < 0 THEN
            RAISE EXCEPTION 'Le prix doit être positif ou nul. Valeur reçue: %', prix;
        END IF;
        
        IF prix > 100000 THEN
            RAISE EXCEPTION 'Le prix ne peut pas dépasser 100000';
        END IF;
        
        -- Validation du nombre de places
        IF nombre_places IS NULL OR nombre_places <= 0 THEN
            RAISE EXCEPTION 'Le nombre de places doit être positif. Valeur reçue: %', nombre_places;
        END IF;
        
        IF nombre_places > 100000 THEN
            RAISE EXCEPTION 'Le nombre de places ne peut pas dépasser 100000';
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Création d'un lieu
-- ================================================
CREATE OR REPLACE FUNCTION creer_lieu(
    p_nom VARCHAR(150),
    p_adresse TEXT,
    p_ville VARCHAR(100),
    p_capacite INT
) RETURNS UUID AS $$
DECLARE
    nouveau_lieu_id UUID;
BEGIN
    INSERT INTO lieu (nom, adresse, ville, capacite)
    VALUES (
        p_nom, 
        p_adresse, 
        p_ville, 
        CASE WHEN p_capacite <= 0 THEN NULL ELSE p_capacite END
    )
    RETURNING id INTO nouveau_lieu_id;
    
    RETURN nouveau_lieu_id;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Création d'un événement
-- ================================================
CREATE OR REPLACE FUNCTION creer_evenement(
    p_titre VARCHAR(150),
    p_description TEXT,
    p_date_debut TIMESTAMP,
    p_date_fin TIMESTAMP,
    p_type_id UUID,
    p_lieu_id UUID
) RETURNS UUID AS $$
DECLARE
    nouvel_evenement_id UUID;
BEGIN
    INSERT INTO evenement (titre, description, date_debut, date_fin, type_id, lieu_id)
    VALUES (p_titre, p_description, p_date_debut, p_date_fin, p_type_id, p_lieu_id)
    RETURNING id INTO nouvel_evenement_id;
    
    RETURN nouvel_evenement_id;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Création des tarifs pour un événement
-- ================================================
CREATE OR REPLACE FUNCTION creer_tarifs(
    p_evenement_id UUID,
    p_tarifs JSONB
) RETURNS VOID AS $$
DECLARE
    tarif_record JSONB;
BEGIN
    FOR tarif_record IN SELECT * FROM jsonb_array_elements(p_tarifs) 
    LOOP
        INSERT INTO tarif (prix, nombre_places, evenement_id, type_place_id)
        VALUES (
            (tarif_record->>'prix')::DECIMAL,
            (tarif_record->>'nombre_places')::INTEGER,
            p_evenement_id,
            (tarif_record->>'type_place_id')::UUID
        );
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Création des fichiers pour un événement
-- ================================================
CREATE OR REPLACE FUNCTION creer_fichiers(
    p_evenement_id UUID,
    p_fichiers JSONB
) RETURNS VOID AS $$
DECLARE
    fichier_record JSONB;
    donnees_binaires BYTEA;
    taille_calculee BIGINT;
BEGIN
    FOR fichier_record IN SELECT * FROM jsonb_array_elements(p_fichiers) 
    LOOP
        -- Conversion base64 vers BYTEA avec gestion d'erreur
        BEGIN
            donnees_binaires := decode(
                REPLACE(REPLACE(fichier_record->>'donnees_binaire', ' ', '+'), '\n', ''), 
                'base64'
            );
        EXCEPTION
            WHEN others THEN
                RAISE EXCEPTION 'Données base64 invalides pour le fichier: %', 
                    fichier_record->>'nom_fichier';
        END;
        
        -- Calcul de la taille
        taille_calculee := octet_length(donnees_binaires);
        
        -- Validation de la taille
        IF taille_calculee = 0 THEN
            RAISE EXCEPTION 'Le fichier % est vide', fichier_record->>'nom_fichier';
        END IF;
        
        IF taille_calculee > 10485760 THEN -- 10MB max
            RAISE EXCEPTION 'Le fichier % dépasse la taille maximale de 10MB', 
                fichier_record->>'nom_fichier';
        END IF;
        
        -- Validation du type de fichier
        IF (fichier_record->>'type_fichier') NOT IN ('photo', 'affiche', 'document') THEN
            RAISE EXCEPTION 'Type de fichier invalide: %', fichier_record->>'type_fichier';
        END IF;
        
        -- Insertion du fichier
        INSERT INTO fichier_evenement (
            evenement_id, 
            nom_fichier, 
            type_mime, 
            taille_bytes, 
            type_fichier, 
            donnees_binaire
        )
        VALUES (
            p_evenement_id,
            fichier_record->>'nom_fichier',
            fichier_record->>'type_mime',
            taille_calculee,
            fichier_record->>'type_fichier',
            donnees_binaires
        );
    END LOOP;
END;
$$ LANGUAGE plpgsql;


-- ================================================
-- FONCTION PRINCIPALE : Création complète d'événement
-- ================================================
CREATE OR REPLACE FUNCTION creer_evenement_complet(
    p_titre VARCHAR(150),
    p_description TEXT,
    p_date_debut TIMESTAMP,
    p_date_fin TIMESTAMP,
    p_type_id UUID,
    p_lieu_nom VARCHAR(150),
    p_lieu_adresse TEXT,
    p_lieu_ville VARCHAR(100),
    p_lieu_capacite INT,
    p_tarifs JSONB,
    p_fichiers JSONB DEFAULT '[]'::JSONB
) 
RETURNS UUID
AS $$
DECLARE
    nouveau_lieu_id UUID;
    nouvel_evenement_id UUID;
    total_places_demandees INTEGER := 0;
BEGIN
    -- ÉTAPE 1: Validation des paramètres
    PERFORM valider_parametres_creation(
        p_titre, p_description, p_date_debut, p_date_fin,
        p_type_id, p_tarifs
    );
    
    -- ÉTAPE 2: Vérification capacité
    total_places_demandees := calculer_total_places(p_tarifs);
    
    IF p_lieu_capacite IS NOT NULL AND total_places_demandees > p_lieu_capacite THEN
        RAISE EXCEPTION 
            'Capacité du lieu insuffisante. Places demandées: %, Capacité: %', 
            total_places_demandees, p_lieu_capacite;
    END IF;
    
    -- ÉTAPE 3: Création du lieu
    nouveau_lieu_id := creer_lieu(
        p_lieu_nom, p_lieu_adresse, p_lieu_ville, p_lieu_capacite
    );
    
    -- ÉTAPE 4: Création de l'événement
    nouvel_evenement_id := creer_evenement(
        p_titre, p_description, p_date_debut, p_date_fin,
        p_type_id, nouveau_lieu_id
    );
    
    -- ÉTAPE 5: Création des tarifs (déclenche automatiquement la création des places)
    PERFORM creer_tarifs(nouvel_evenement_id, p_tarifs);
    
    -- ÉTAPE 6: Création des fichiers (optionnel)
    IF p_fichiers IS NOT NULL AND jsonb_array_length(p_fichiers) > 0 THEN
        PERFORM creer_fichiers_bytea(nouvel_evenement_id, p_fichiers);
    END IF;
    
    RETURN nouvel_evenement_id;
END;
$$ LANGUAGE plpgsql;


-- ================================================
-- TRIGGER : Validation création/mise à jour d'événement
-- ================================================
CREATE OR REPLACE FUNCTION valider_nouvel_evenement()
RETURNS TRIGGER AS $$
BEGIN
    -- 1. Vérification que la date de début est dans le futur (pour INSERT seulement)
    IF TG_OP = 'INSERT' AND NEW.date_debut <= CURRENT_TIMESTAMP THEN
        RAISE EXCEPTION 'La date de début doit être dans le futur';
    END IF;
    
    -- 2. Vérification que la date de fin est après la date de début
    IF NEW.date_debut >= NEW.date_fin THEN
        RAISE EXCEPTION 'La date de début doit être avant la date de fin';
    END IF;
    
    -- 3. Vérification des chevauchements de réservation du lieu
    IF EXISTS (
        SELECT 1 FROM evenement e
        WHERE e.lieu_id = NEW.lieu_id
        AND e.id != NEW.id
        AND e.date_debut < NEW.date_fin
        AND e.date_fin > NEW.date_debut
    ) THEN
        RAISE EXCEPTION 'Le lieu est déjà réservé sur cette plage horaire';
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_valider_evenement
    BEFORE INSERT OR UPDATE ON evenement
    FOR EACH ROW
    EXECUTE FUNCTION valider_nouvel_evenement();

-- ================================================
-- TRIGGER : Création automatique des places
-- ================================================
CREATE OR REPLACE FUNCTION creer_places_automatiquement()
RETURNS TRIGGER AS $$
DECLARE
    compteur INTEGER := 1;
    numero_place VARCHAR(100);
    nom_type_place VARCHAR(50);
    uuid_evenement TEXT;
BEGIN
    -- Récupérer le nom du type de place et l'UUID de l'événement
    SELECT tp.nom, e.id
    INTO nom_type_place, uuid_evenement
    FROM type_place tp
    JOIN evenement e ON e.id = NEW.evenement_id
    WHERE tp.id = NEW.type_place_id;
    
    -- Vérifier que les données existent
    IF nom_type_place IS NULL OR uuid_evenement IS NULL THEN
        RAISE EXCEPTION 'Type de place ou événement non trouvé';
    END IF;

    -- Création des places avec le format demandé
    WHILE compteur <= NEW.nombre_places LOOP
        -- Format: TypePlace-UUIDEvenement-NuméroAuto
        numero_place := CONCAT(
            UPPER(nom_type_place),
            '-', 
            uuid_evenement,
            '-', 
            LPAD(compteur::TEXT, 3, '0')
        );
        
        -- Insérer la place
        INSERT INTO place (numero, etat_code, tarif_id)
        VALUES (numero_place, 'disponible', NEW.id);
        
        compteur := compteur + 1;
    END LOOP;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_creer_places
    AFTER INSERT ON tarif
    FOR EACH ROW
    EXECUTE FUNCTION creer_places_automatiquement();

-- ================================================
-- TRIGGER : Audit des changements d'état des places
-- ================================================
CREATE OR REPLACE FUNCTION auditer_changement_etat()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.etat_code IS DISTINCT FROM NEW.etat_code THEN
        INSERT INTO audit_place (place_id, ancien_etat, nouvel_etat, utilisateur)
        VALUES (NEW.id, OLD.etat_code, NEW.etat_code, CURRENT_USER);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_audit_place
    AFTER UPDATE ON place
    FOR EACH ROW
    EXECUTE FUNCTION auditer_changement_etat();




-- ================================================
-- FONCTION : Vérification de capacité (version simplifiée)
-- ================================================
CREATE OR REPLACE FUNCTION verifier_capacite_lieu(
    p_lieu_capacite INT,
    p_tarifs JSONB
) RETURNS BOOLEAN AS $$
DECLARE
    total_places_demandees INTEGER := 0;
    tarif_record JSONB;
BEGIN
    -- Si capacité NULL = illimité → toujours valide
    IF p_lieu_capacite IS NULL THEN
        RETURN TRUE;
    END IF;
    
    -- Calculer le TOTAL des nouvelles places demandées
    FOR tarif_record IN SELECT * FROM jsonb_array_elements(p_tarifs) 
    LOOP
        total_places_demandees := total_places_demandees + 
            COALESCE((tarif_record->>'nombre_places')::INTEGER, 0);
    END LOOP;
    
    -- Vérifier si la capacité est suffisante
    RETURN total_places_demandees <= p_lieu_capacite;
END;
$$ LANGUAGE plpgsql;

-- ================================================
-- FONCTION : Génération de numéros de place (optionnelle)
-- ================================================
CREATE OR REPLACE FUNCTION generer_numero_place(
    p_type_place_id UUID, 
    p_numero INTEGER
) RETURNS VARCHAR(50) AS $$
DECLARE
    nom_type VARCHAR(50);
BEGIN
    SELECT nom INTO nom_type FROM type_place WHERE id = p_type_place_id;
    RETURN CONCAT(UPPER(SUBSTRING(nom_type FROM 1 FOR 3)), '-', LPAD(p_numero::TEXT, 3, '0'));
END;
$$ LANGUAGE plpgsql;


-- APRES MODIF /

ALTER TABLE place DROP CONSTRAINT IF EXISTS uq_place_numero_tarif;


CREATE OR REPLACE FUNCTION creer_fichiers_bytea(
    p_evenement_id UUID,
    p_fichiers JSONB
) RETURNS VOID AS $$
DECLARE
    fichier_record JSONB;
    donnees_binaires BYTEA;
    taille_calculee BIGINT;
BEGIN
    FOR fichier_record IN SELECT * FROM jsonb_array_elements(p_fichiers) 
    LOOP
        -- ✅ Conversion directe depuis le JSON (supposant que c'est déjà en hex/bytea)
        -- Si vos données sont encodées d'une manière spécifique dans le JSON
        donnees_binaires := decode(fichier_record->>'donnees_bytea', 'hex');
        
        -- Calcul de la taille
        taille_calculee := octet_length(donnees_binaires);
        
        IF taille_calculee = 0 OR taille_calculee IS NULL THEN
            RAISE EXCEPTION 'Le fichier % est vide ou invalide', fichier_record->>'nom_fichier';
        END IF;
        
        INSERT INTO fichier_evenement (
            evenement_id, nom_fichier, type_mime, taille_bytes, type_fichier, donnees_binaire
        )
        VALUES (
            p_evenement_id,
            fichier_record->>'nom_fichier',
            fichier_record->>'type_mime',
            taille_calculee,
            fichier_record->>'type_fichier',
            donnees_binaires
        );
    END LOOP;
END;
$$ LANGUAGE plpgsql;