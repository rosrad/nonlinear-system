function peaks(inf)
    [A,del] = importdata(inf);
    [peaks, locs] = findpeaks(A(:,2)');
    out = A(locs',:)
    dlmwrite(strcat(inf,'.mlt'), out, ' ')
end
