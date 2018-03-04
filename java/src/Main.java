import java.util.*;
import java.io.*;
import java.math.*;

/**
 * Made with love by AntiSquid, Illedan and Wildum.
 * You can help children learn to code while you participate by donating to CoderDojo.
 **/
class Player {
    
    public static class GameState {
        public int myTeam;
        public List<Item> items;
        public List<Item> potions;
        
        public GameState(Scanner in) {
            this.myTeam = in.nextInt();
            
            // TODO: later
            int bushAndSpawnPointCount = in.nextInt(); // usefull from wood1, represents the number of bushes and the number of places where neutral units can spawn
            for (int i = 0; i < bushAndSpawnPointCount; i++) {
                String entityType = in.next(); // BUSH, from wood1 it can also be SPAWN
                int x = in.nextInt();
                int y = in.nextInt();
                int radius = in.nextInt();
            }
            
            int itemCount = in.nextInt(); // useful from wood2
            this.items = new ArrayList<>();
            this.potions = new ArrayList<>();
            for (int i = 0; i < itemCount; i++) {
                Item item = new Item(in);
                if (item.isPotion) {
                    this.potions.add(item);
                } else {
                    this.items.add(item);
                }
            }
        }
    }
    
    public static class RoundState {
        public int gold;
        public int enemyGold;
        public int type;
        
        public List<Entity> allies;
        public List<Entity> enemies;
        public List<Entity> neutrals;
        public List<Hero> heroes;
        public Entity tower;
        
        public RoundState(Scanner in, GameState gs) {
            this.gold = in.nextInt();
            this.enemyGold = in.nextInt();
            this.type = in.nextInt(); // a positive value will show the number of heroes that await a command
            
            int entityCount = in.nextInt();
            
            // Get all useful entities here
            this.allies = new ArrayList<>();
            this.enemies = new ArrayList<>();
            this.neutrals = new ArrayList<>();
            this.heroes = new ArrayList<>();
            this.tower = null;

            for (int i = 0; i < entityCount; i++) {
                Entity entity = new Entity(in);
                if (entity.team == gs.myTeam) { // allies
                    if (entity.unitType.equals("HERO")) {
                        this.heroes.add(new Hero(entity));
                    } else {
                        this.allies.add(entity);
                    }
                    if (entity.unitType.equals("TOWER")) {
                        this.tower = entity;
                    }
                } else if (entity.team == -1) {
                    this.neutrals.add(entity);
                } else {
                    this.enemies.add(entity);
                }
            }
        }
    }
    
    public static class Item {
        public String name;
        public int cost;
        public int damage;
        public int health;
        public int maxHealth;
        public int mana;
        public int maxMana;
        public int moveSpeed;
        public int manaRegeneration;
        public boolean isPotion;
        
        public Item(Scanner in) {
            this.name = in.next(); // contains keywords such as BRONZE, SILVER and BLADE, BOOTS connected by "_" to help you sort easier
            this.cost = in.nextInt(); // BRONZE items have lowest cost, the most expensive items are LEGENDARY
            this.damage = in.nextInt(); // keyword BLADE is present if the most important item stat is damage
            this.health = in.nextInt();
            this.maxHealth = in.nextInt();
            this.mana = in.nextInt();
            this.maxMana = in.nextInt();
            this.moveSpeed = in.nextInt(); // keyword BOOTS is present if the most important item stat is moveSpeed
            this.manaRegeneration = in.nextInt();
            this.isPotion = in.nextInt() > 0; // 0 if it's not instantly consumed
        }
        
        public String toString() {
            String s = "Name: " + this.name +
                       ", Cost: " + this.cost;
                       
            if (this.damage > 0) s += ", Damage: " + this.damage;
            if (this.health > 0) s += ", Health: " + this.health;
            if (this.maxHealth > 0) s += ", MaxHealth: " + this.maxHealth;
            if (this.mana > 0) s += ", Mana: " + this.mana;
            if (this.maxMana > 0) s += ", MaxMana: " + this.maxMana;
            if (this.moveSpeed > 0)  s += ", MoveSpeed: " + this.moveSpeed;
            if (this.manaRegeneration > 0) s += ", ManaRegeneration: " + this.manaRegeneration;
            if (this.isPotion) s += ", IsPotion";
            
            return s;
        }
    }
    
    public static class Entity {
        public int unitId;
        public int team;
        public String unitType;
        public int x;
        public int y;
        public int attackRange;
        public int health;
        public int maxHealth;
        public int shield;
        public int attackDamage;
        public int movementSpeed;
        public int stunDuration;
        public int goldValue;
        
        public double distanceToHero;
        public double distanceToTower;
        
        // HERO
        public int countDown1;
        public int countDown2;
        public int countDown3;
        public int mana;
        public int maxMana;
        public int manaRegeneration;
        public String heroType;
        public int isVisible;
        public int itemsOwned;
        
        public Entity(Scanner in) {
            this.unitId = in.nextInt();
            this.team = in.nextInt();
            this.unitType = in.next();
            this.x = in.nextInt();
            this.y = in.nextInt();
            this.attackRange = in.nextInt();
            this.health = in.nextInt();
            this.maxHealth = in.nextInt();
            this.shield = in.nextInt();
            this.attackDamage = in.nextInt();
            this.movementSpeed = in.nextInt();
            this.stunDuration = in.nextInt();
            this.goldValue = in.nextInt();
            
            // HERO
            this.countDown1 = in.nextInt(); // all countDown and mana variables are useful starting in bronze
            this.countDown2 = in.nextInt();
            this.countDown3 = in.nextInt();
            this.mana = in.nextInt();
            this.maxMana = in.nextInt();
            this.manaRegeneration = in.nextInt();
            this.heroType = in.next(); // DEADPOOL, VALKYRIE, DOCTOR_STRANGE, HULK, IRONMAN
            this.isVisible = in.nextInt(); // 0 if it isn't
            this.itemsOwned = in.nextInt(); // useful from wood1
        }
        
        public String toString() {
            return "ID: " + this.unitId + ", Type: " + this.unitType + ", Team: " + this.team + ", X: " + this.x + ", Y: " + this.y;
        }
        
        // distanceTo returns the distance between this entity and another.
        public double distanceTo(Entity e) {
            return Math.abs(Math.sqrt(Math.pow(e.x - this.x, 2) + Math.pow(e.y - this.y, 2)));
        }
        
        public void attack(int unitId) {
            System.out.printf("ATTACK %d\n", unitId);
        }
        
        public void distanceToMove(double distanceX, double distanceY) {
            if (this.team == 0) {
                System.out.printf("MOVE %f %f\n", this.x+distanceX, this.y+distanceY);
            } else {
                System.out.printf("MOVE %f %f\n", this.x-distanceX, this.y-distanceY);
            }
        }
        
        public void moveAndAttack(double distanceX, double distanceY, int unitId) {
            if (this.team == 0) {
                System.out.printf("MOVE_ATTACK %f %f %d\n", this.x+distanceX, this.y+distanceY, unitId);
            } else {
                System.out.printf("MOVE_ATTACK %f %f %d\n", this.x-distanceX, this.y-distanceY, unitId);
            }
        }
    }
    
    public static class Hero {
        public Entity entity;
        public HeroInventory inventory;
        
        public Entity farestAlly;
        public Entity nearestEnemy;
        
        public Entity lastHitableEnemy;
        public Entity lastHitableAlly;
        
        public Hero(Entity entity) {
            this.entity = entity;
        }
        
        public boolean Defense(GameState gs, RoundState rs) {
            this.farestAlly = null;
            this.lastHitableAlly = null;
            for (Entity ally : rs.allies) {
                ally.distanceToHero = this.entity.distanceTo(ally);
                ally.distanceToTower = rs.tower.distanceTo(ally);
                
                // the farest ally is the farest from the tower.
                if (this.farestAlly == null || ally.distanceToTower >= this.farestAlly.distanceToTower) {
                    this.farestAlly = ally;
                }

                // loop on all allies to find ally on hero range
                if (ally.distanceToHero < this.entity.attackRange + this.entity.movementSpeed) {
                    // find lastHitable ally
                    if (ally.health <= this.entity.attackDamage) {
                        this.lastHitableAlly = ally;
                    }
                }
            }
            
            this.nearestEnemy = null;
            for (Entity bad : rs.enemies) {
                bad.distanceToHero = this.entity.distanceTo(bad);
                bad.distanceToTower = rs.tower.distanceTo(bad);

                // find the nearest enemy with tower distance
                if (this.nearestEnemy == null || bad.distanceToTower < this.nearestEnemy.distanceToTower) {
                    this.nearestEnemy = bad;
                }
            }


            if (this.entity.distanceTo(rs.tower) < 100) {
                return false;
            }
            
            int safeDistance = this.entity.movementSpeed / 2;
            // if the farest ally is farest from our tower than the hero
            if (this.farestAlly.distanceToTower < rs.tower.distanceTo(this.entity)) {
                // if we got no ally in front of us, move back
                if (this.farestAlly.distanceToHero > this.entity.movementSpeed) {
                    this.entity.distanceToMove(-this.entity.movementSpeed, 0);

                } else if (this.farestAlly.distanceToHero + this.nearestEnemy.distanceToHero + safeDistance < this.entity.attackRange) {
                    this.entity.moveAndAttack(-this.farestAlly.distanceToHero-safeDistance, 0, this.nearestEnemy.unitId);

                } else {
                    this.entity.distanceToMove(-this.farestAlly.distanceToHero-safeDistance, 0);

                }
                return true;
            }
            
            return false;
        }
        
        public boolean Purchase(GameState gs, RoundState rs) {
            Item bestPotion = null;
            for (Item potion : gs.potions) {
                if (potion.health > 0 // is a health potion
                    && this.entity.maxHealth - this.entity.health >= potion.health // hero lost more than potion health
                    && rs.gold >= potion.cost) { // can afford potion
                    bestPotion = potion;
                }
            }
            
            if (bestPotion != null) {
                System.out.printf("BUY %s\n", bestPotion.name);
                rs.gold -= bestPotion.cost;
                return true;
            }
            
            // Purchase decisions
            Item bestItem = null;
            for (Item item : gs.items) {
                if (item.damage > 0 && rs.gold >= item.cost) {
                    if (bestItem == null || bestItem.damage < item.damage) {
                        bestItem = item;
                    }
                }
            }
            
            if (bestItem != null) {
                // if we got space on bag, buy item
                if (this.entity.itemsOwned < 3) {
                    this.inventory.buy(bestItem);
                    rs.gold -= bestItem.cost;
                    return true;
                } else {
                    // check if item is better than any of hero item.
                    System.err.println(inventory.weakestItem);
                    if (this.inventory.weakestItem != null && bestItem.damage > this.inventory.weakestItem.damage) {
                        this.inventory.sell(this.inventory.weakestItem);
                        return true;
                    }
                }
            }
            
            return false;
        }
        
        public boolean Skill(GameState gs, RoundState rs) {
            switch (this.entity.heroType) {
                case "IRONMAN":
                    // try cast fireball, skill 2
                    for (Entity enemy : rs.enemies) {
                            if (enemy.distanceTo(this.entity) < 250 &&
                                this.entity.mana >= 50 &&
                                this.entity.countDown3 == 0) {
                                System.out.printf("BURNING %d %d\n", enemy.x, enemy.y);
                                return true;
                            }
                    }
                    break;
                case "DOCTOR_STRANGE":
                    for (Entity enemy : rs.enemies) {
                        if (enemy.unitType.equals("HERO")) {
                            if (enemy.distanceTo(this.entity) < 300 &&
                                this.entity.mana >= 40 &&
                                this.entity.countDown3 == 0 &&
                                this.entity.distanceTo(rs.tower) < rs.tower.attackRange - 200) {
                                System.out.printf("PULL %d\n", enemy.unitId);
                                return true;
                            }
                        }
                    }
                    break;
            }
            
            return false;
        }
        
        public boolean Attack(GameState gs, RoundState rs) {
            Entity lastHitable = null;
            for (Entity bad : rs.enemies) {
                // if the distance of the closest enemy to the hero is lower than the hero range attack +
                // hero movementSpees minus the distance of the closest ally of the hero minus the distance from this sbire to the hero
                if (bad.distanceToHero < (this.entity.attackRange + this.entity.movementSpeed) - (int)(Math.floor(this.farestAlly.distanceToHero))) {

                    // if enemy is in attack range of hero, check if enemy is hero.
                    // if enemy is hero, target it, otherwise check if it is a last hitable enemy and if it safe to attack hero.
                    // if nearest enemy is closer to the tower than the farest ally, do not target hero.
                    if (bad.unitType.equals("HERO") && nearestEnemy.distanceToTower < farestAlly.distanceToTower) {
                        System.err.println("ATTACK HERO");
                        if (nearestEnemy == null || bad.health < nearestEnemy.health) {
                            nearestEnemy = bad;
                        }
                        break;

                    } else if (bad.health <= this.entity.attackDamage && (lastHitable == null || bad.health < lastHitable.health)) {
                        lastHitable = bad;

                    }
                }
            }
            
            Entity target = nearestEnemy;
            // if we got a lastHitable enemy
            if (lastHitable != null) {
                System.err.println("LAST HIT");
                target = lastHitable;
                rs.enemies.remove(target);

            } else if (lastHitableAlly != null) {
                System.err.println("DENY");
                target = lastHitableAlly;

            }
            
            System.err.println(target);
            int heroRange = this.entity.attackRange;
            boolean didAttack = false;

            if (heroRange >= target.distanceToHero) { // inside hero range
                if (heroRange * 0.9 < target.distanceToHero) {
                    this.entity.attack(target.unitId);

                } else {
                    this.entity.moveAndAttack(- (heroRange * 0.9 - target.distanceToHero), 0, target.unitId);

                }
                didAttack = true;
            } else { // outside hero range
                if (heroRange+this.entity.movementSpeed > target.distanceToHero) { // if enemy is inside movement range + attack range
                    // move and attack
                    this.entity.moveAndAttack(target.distanceToHero - heroRange+10, 0, target.unitId);
                    didAttack = true;

                } else {
                    // move
                    this.entity.distanceToMove(this.entity.movementSpeed, 0);

                }
            }
            
            if (didAttack) {
                if (target.team == gs.myTeam) {
                    rs.allies.remove(target);

                } else {
                    rs.enemies.remove(target);

                }
            }
            
            return true;
        }
    }
    
    public static class HeroInventory {
        public List<Item> items;
        public Item weakestItem;
        
        public HeroInventory() {
            this.items = new ArrayList<Item>();
        }
        
        public void buy(Item item) {
            System.out.printf("BUY %s\n", item.name);
            this.items.add(item); // add item to hero items
            this.computeWeakestItem();
        }
        
        public void sell(Item item) {
            System.out.printf("SELL %s\n", item.name);
            this.items.remove(item); // remove item from hero items
            this.computeWeakestItem();
        }
        
        private void computeWeakestItem() {
            // compute weakest item
            this.weakestItem = null;
            for (Item item : this.items) {
                if (this.weakestItem == null || item.damage < this.weakestItem.damage) {
                    this.weakestItem = item;
                }
            }
        }
    }

    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        
        GameState gs = new GameState(in);

        // game loop
        List<String> heroPicks = Arrays.asList("IRONMAN", "HULK");
        int pick = 0;
        
        Map<String, HeroInventory> heroesInventories = new HashMap<String, HeroInventory>();
        for (String heroPick : heroPicks) {
            heroesInventories.put(heroPick, new HeroInventory());
        }
        
        while (true) {
            RoundState rs = new RoundState(in, gs);
            
            // Write an action using System.out.println()
            // To debug: System.err.println("Debug messages...");
            // Hero picks
            if (rs.type < 0) {
                System.out.println(heroPicks.get(pick++));
                continue;
            }

            // If roundType has a negative value then you need to output a Hero name, such as "DEADPOOL" or "VALKYRIE".
            // Else you need to output roundType number of any valid action, such as "WAIT" or "ATTACK unitId"
            for (Hero hero : rs.heroes) {
                hero.inventory = heroesInventories.get(hero.entity.heroType);
                
                // Defense section
                if (hero.Defense(gs, rs)) {
                    System.err.printf("HERO %s : DEFENSE\n", hero.entity.heroType);
                    continue;
                }
                
                // Purchase section
                if (hero.Purchase(gs, rs)) {
                    System.err.printf("HERO %s : PURCHASE\n", hero.entity.heroType);
                    continue;
                }
                
                // Skill section
                if (hero.Skill(gs, rs)) {
                    System.err.printf("HERO %s : SKILL\n", hero.entity.heroType);
                    continue;
                }
                
                // Attack section
                // find the nearest enemy
                if (hero.Attack(gs, rs)) {
                    System.err.printf("HERO %s : ATTACK\n", hero.entity.heroType);
                    continue;
                }
            }
        }
    }
}