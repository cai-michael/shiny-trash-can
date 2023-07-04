% All details from the story to generate our answer space.
name(winslow; marcolla; contee; natsiou; finch).
color(white; green; purple; red; blue).
drink(whiskey; absinthe; wine; rum; beer).
chair(1..5).
country(dunwall; dabokva; baleton; karnaca; fraeport).
trinket(birdpendant; ring; snufftin; diamond; medal).

% Create distinct sets of characteristics, no one can share any individual one.
1 {details(A,B,C,D,E,F) : color(B), drink(C), chair(D), country(E), trinket(F)} 1 :- name(A).
1 {details(A,B,C,D,E,F) : name(A), drink(C), chair(D), country(E), trinket(F)} 1 :- color(B).
1 {details(A,B,C,D,E,F) : name(A), color(B), chair(D), country(E), trinket(F)} 1 :- drink(C).
1 {details(A,B,C,D,E,F) : name(A), color(B), drink(C), country(E), trinket(F)} 1 :- chair(D).
1 {details(A,B,C,D,E,F) : name(A), color(B), drink(C), chair(D), trinket(F)} 1 :- country(E).
1 {details(A,B,C,D,E,F) : name(A), color(B), drink(C), chair(D), country(E)} 1 :- trinket(F).

% Interesting note: since an implication with nothing on the left means the statement on the right should always be false we need to boolean not a lot of these, weird :)

% Madam Natsiou wore a jaunty green hat.
:- not details(natsiou,green,_,_,_,_).

% Doctor Marcolla was at the far left... next to the guest wearing a purple jacket.
:- not details(marcolla,_,_,1,_,_).
:- not details(_,purple,_,2,_,_).

% The lady in red sat left of someone in blue.
:- details(_,red,_,A,_,_), details(_,blue,_,B,_,_), B-A!=1.

% I remember that red outfit because the woman spilled her whiskey all over it.
:- not details(_,red,whiskey,_,_,_).

% The traveler from Dunwall was dressed entirely in white.
:- not details(_,white,_,_,dunwall,_).

% When one of the dinner guests bragged about her Diamond, the woman next to her said they were finer in Dunwall, where she lived.
:- details(_,_,_,_,dunwall,diamond).
:- details(_,_,_,A,dunwall,_), details(_,_,_,B,_,diamond), |A-B|!=1.

% So Lady Winslow showed off a prized War Medal, at which the lady from Fraeport scoffed, saying it was no match for her Ring.
:- not details(winslow,_,_,_,_,medal).
:- details(winslow,_,_,_,fraeport,_).
:- not details(_,_,_,_,fraeport,ring).

% Someone else carried a valuable Snuff Tin and when she saw it, the visitor from Karnaca next to her almost spilled her neighbor's rum.
:- details(_,_,_,_,karnaca,snufftin).
:- details(_,_,_,A,karnaca,_), details(_,_,rum,B,_,_), |A-B|!=1.
:- details(_,_,_,A,karnaca,_), details(_,_,_,B,_,snufftin), |A-B|!=1.

% Countess Contee raised her beer in toast.
:- not details(contee,_,beer,_,_,_).

% The lady from Baleton, full of wine, jumped up onto the table, falling onto the guest in the center seat, spilling the poor woman's absinthe.
:- not details(_,_,wine,_,baleton,_).
:- not details(_,_,absinthe,3,_,_).

% Then Baroness Finch captivated them all with a story about her wild youth in Dabokva.
:- not details(finch,_,_,_,dabokva,_).